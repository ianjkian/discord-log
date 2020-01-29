// Copyright (c) 2020 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package discordlog

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type client struct {
	apiId       string
	apiToken    string
	apiDomain   string
	apiWebhook  string
	httpClient  *http.Client
	httpTimeout time.Duration
}

// newClient returns an instantiated Client struct.
func newClient(apiId, apiToken string) (c *client, err error) {
	if apiId == "" {
		apiId = os.Getenv(envId)
	}
	if apiToken == "" {
		apiToken = os.Getenv(envToken)
	}
	c = &client{
		apiId:       apiId,
		apiToken:    apiToken,
		apiDomain:   apiDomain,
		apiWebhook:  apiWebhook,
		httpClient:  &http.Client{},
		httpTimeout: apiRequestTimeout * time.Second,
	}
	return
}

// do returns response from HTTP request or returns error if time exceeded.
func (c *client) do(request *http.Request) (*http.Response, error) {
	type reqres struct {
		response *http.Response
		err      error
	}
	timer := time.NewTimer(c.httpTimeout)
	done := make(chan reqres, 1)
	go func() {
		r, err := c.httpClient.Do(request)
		done <- reqres{r, err}
	}()
	select {
	case r := <-done:
		return r.response, r.err
	case <-timer.C:
		return nil, errors.New(msgRequestTimeout)
	}
}

// call prepares and process HTTP request to endpoint.
func (c *client) call(payload []byte) (result []byte, err error) {
	var req *http.Request
	var resp *http.Response
	rawurl := fmt.Sprintf("%s/%s/%s/%s", c.apiDomain, c.apiWebhook, c.apiId, c.apiToken)
	if req, err = http.NewRequest("POST", rawurl, bytes.NewReader(payload)); err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if resp, err = c.do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	if result, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	return
}

func (c *client) handleRequest(payload interface{}) (json.RawMessage, error) {
	var js, rawresponse []byte
	var err error
	if js, err = json.Marshal(payload); err != nil {
		return nil, err
	}
	if rawresponse, err = c.call(js); err != nil {
		return nil, err
	}
	return rawresponse, nil
}

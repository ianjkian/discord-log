// Copyright (c) 2020 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package discordlog

import (
	"time"
)

type embed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Timestamp   string `json:"timestamp"`
}

type embeds []embed

// PrintInfo sends info message to webhook.
func (wh *Webhook) PrintInfo(text string) error {
	e := embed{
		Title:       "Info",
		Description: text,
		Color:       colorBlack,
		Timestamp:   time.Now().Format(time.RFC3339),
	}
	err := wh.printEmbed(e)
	return err
}

// PrintWarning sends warning message to webhook.
func (wh *Webhook) PrintWarning(text string) error {
	e := embed{
		Title:       "Warning",
		Description: text,
		Color:       colorYellow,
		Timestamp:   time.Now().Format(time.RFC3339),
	}
	err := wh.printEmbed(e)
	return err
}

// PrintError sends error message to webhook.
func (wh *Webhook) PrintError(text string) error {
	e := embed{
		Title:       "Error",
		Description: text,
		Color:       colorRed,
		Timestamp:   time.Now().Format(time.RFC3339),
	}
	err := wh.printEmbed(e)
	return err
}

func (wh *Webhook) printEmbed(e embed) error {
	c, err := newClient(wh.id, wh.token)
	if err != nil {
		return err
	}
	q := map[string]embeds{
		"embeds": embeds{e},
	}
	_, err = c.handleRequest(q)
	return err
}

// Copyright (c) 2020 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package discordlog

// Webhook struct
type Webhook struct {
	id    string
	token string
}

// NewWebhook returns an instantiated Webhook struct.
func NewWebhook(id, token string) *Webhook {
	wh := &Webhook{id: id, token: token}
	return wh
}

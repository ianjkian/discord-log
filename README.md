discord-log [![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/nikchis/discord-log/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/nikchis/discord-log?)](https://goreportcard.com/report/github.com/nikchis/discord-log) [![GoDoc](https://godoc.org/github.com/nikchis/discord-log?status.svg)](https://godoc.org/github.com/nikchis/discord-log)
==========

discord-log is a golang package which allows you easily send text messages with different levels of logging {info, warning, error} to Discord webhook.

## Import
	import "github.com/nikchis/discord-log"
	
## Usage

~~~ go
package main

import discordlog "github.com/nikchis/discord-log"

const (
	WebhookId    = "YOUR_WEBHOOK_ID"
	WebhookToken = "YOUR_WEBHOOK_TOKEN"
)

func main() {
	dlog := discordlog.NewWebhook(WebhookId, WebhookToken)
	// print info, warning, error messages
	dlog.PrintInfo("Hello! This is info message.")
	dlog.PrintWarning("Attention! This is warning message.")
	dlog.PrintError("Oops! Error occured! This is error message.")
}
~~~

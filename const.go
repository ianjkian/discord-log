// Copyright (c) 2020 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package discordlog

const (
	apiDomain         = "https://discordapp.com/api"
	apiWebhook        = "webhooks"
	apiRequestTimeout = 30

	colorRed    = "13369344"
	colorYellow = "16763904"
	colorBlack  = "1118481"

	envId    = "DISCORD_WEBHOOK_ID"
	envToken = "DISCORD_WEBHOOK_TOKEN"

	msgEmptyArgs      = "Empty args"
	msgRequestTimeout = "Request timeout"
	msgUnsupType      = "Unsupported arg type"
)

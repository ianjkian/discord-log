// Copyright (c) 2020 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license

package discordlog

import (
	"testing"
)

const (
	Id    = ``
	Token = ``
)

func TestPrintWarning(t *testing.T) {
	dlog := NewWebhook(Id, Token)
	if err := dlog.PrintWarning("new warning occurred!"); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestPrintError(t *testing.T) {
	dlog := NewWebhook(Id, Token)
	if err := dlog.PrintError("new error occurred!"); err != nil {
		t.Error(err)
		t.Fail()
	}
}

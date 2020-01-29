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
	if err := PrintWarning("new warning occured!", Id, Token); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestPrintError(t *testing.T) {
	if err := PrintError("new error occured!", Id, Token); err != nil {
		t.Error(err)
		t.Fail()
	}
}

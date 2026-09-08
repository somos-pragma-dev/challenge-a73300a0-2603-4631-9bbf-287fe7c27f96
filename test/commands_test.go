package test

import (
	"testing"
	"internal/commands"
	"github.com/stretchr/testify/assert"
)

func TestBalanceCmd(t *testing.T) {
	cmd := commands.BalanceCmd()
	assert.NotNil(t, cmd)
}

func TestTransferCmd(t *testing.T) {
	cmd := commands.TransferCmd()
	assert.NotNil(t, cmd)
}
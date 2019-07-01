package xjson

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFile(t *testing.T) {
	var me struct {
		Having string `json:"having"`
	}
	require.NoError(t, ParseFile("test.json", &me))
	require.Equal(t, "coffee", me.Having)
}

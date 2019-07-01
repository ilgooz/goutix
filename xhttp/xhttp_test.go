package xhttp

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHasHeader(t *testing.T) {
	r := &http.Request{
		Header: http.Header{"x": []string{"1", "2"}},
	}
	require.True(t, HasHeader(r, "x", nil))
	require.False(t, HasHeader(r, "y", nil))
	require.True(t, HasHeader(r, "x", []string{"1"}))
	require.True(t, HasHeader(r, "x", []string{"1", "2"}))
	require.False(t, HasHeader(r, "x", []string{"1", "3"}))
}

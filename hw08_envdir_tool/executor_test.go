package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("bad cmd", func(t *testing.T) {
		code := RunCmd([]string{}, Environment{})
		require.Equalf(t, 1, code, "should be error code 1")
	})
}

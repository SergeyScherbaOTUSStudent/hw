package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListItem(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		i := newListItem("Val", "K")

		require.Equal(t, "Val", i.Value)
		require.Equal(t, Key("K"), i.Key)
		require.Nil(t, i.Prev)
		require.Nil(t, i.Next)
	})
}

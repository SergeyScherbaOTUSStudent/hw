package hw04lrucache

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListItem(t *testing.T) {
	t.Run("test ListItem constructor", func(t *testing.T) {
		i := newListItem("Val")

		require.Equal(t, "Val", i.Value)
		require.Nil(t, i.Prev)
		require.Nil(t, i.Next)
	})
}

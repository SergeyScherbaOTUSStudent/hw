package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListItem(t *testing.T) {
	t.Run("test ListItem constructor", func(t *testing.T) {
		i := newListItem("Val")

		require.Equal(t, "Val", i.Value)
		require.Nil(t, i.Prev)
		require.Nil(t, i.Next)
	})
}

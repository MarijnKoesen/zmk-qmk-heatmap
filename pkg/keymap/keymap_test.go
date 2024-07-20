package keymap

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeymap(t *testing.T) {
	m := New()
	require.Equal(t, 0, m.NumberOfKeys())

	m.AddLayer("default", []Row{
		{Keys: []Key{{Tap: "Q"}, {Tap: "W"}}},
		{Keys: []Key{{Tap: "A"}, {Tap: "R"}}},
		{Keys: []Key{{Tap: "SPACE"}}},
	})

	require.Equal(t, 5, m.NumberOfKeys())
}

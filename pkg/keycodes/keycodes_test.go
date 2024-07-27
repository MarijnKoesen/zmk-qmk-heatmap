package keycodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var keyCodeTests = []struct {
	key             KeyCode
	expectedKeyCode int16
}{
	{KC_A, 0x04},
	{KC_KP_A, 0xBC},
	{KC_LCTRL, 0xE0},
}

func TestKeyCodes(t *testing.T) {
	for _, test := range keyCodeTests {
		assert.Equal(t, KeyCode(test.expectedKeyCode), test.key)
	}
}

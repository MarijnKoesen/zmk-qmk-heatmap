package keycodes

type KeyCode int

/* USB HID Keyboard/Keypad Usage(0x07) */
const (
	KC_NO             KeyCode = iota
	KC_ROLL_OVER              // 01
	KC_POST_FAIL              // 02
	KC_UNDEFINED              // 03
	KC_A                      // 04
	KC_B                      // 05
	KC_C                      // 06
	KC_D                      // 07
	KC_E                      // 08
	KC_F                      // 09
	KC_G                      // 0A
	KC_H                      // 0B
	KC_I                      // 0C
	KC_J                      // 0D
	KC_K                      // 0E
	KC_L                      // 0F
	KC_M                      // 10
	KC_N                      // 11
	KC_O                      // 12
	KC_P                      // 13
	KC_Q                      // 14
	KC_R                      // 15
	KC_S                      // 16
	KC_T                      // 17
	KC_U                      // 18
	KC_V                      // 19
	KC_W                      // 1A
	KC_X                      // 1B
	KC_Y                      // 1C
	KC_Z                      // 1D
	KC_1                      // 1E
	KC_2                      // 1F
	KC_3                      // 20
	KC_4                      // 21
	KC_5                      // 22
	KC_6                      // 23
	KC_7                      // 24
	KC_8                      // 25
	KC_9                      // 26
	KC_0                      // 27
	KC_ENTER                  // 28
	KC_ESCAPE                 // 29
	KC_BSPACE                 // 2A
	KC_TAB                    // 2B
	KC_SPACE                  // 2C
	KC_MINUS                  // 2D
	KC_EQUAL                  // 2E
	KC_LBRACKET               // 2F
	KC_RBRACKET               // 30
	KC_BSLASH                 // 31   \ (and |)
	KC_NONUS_HASH             // 32   Non-US # and ~ (Typically near the Enter key)
	KC_SCOLON                 // 33   ; (and :)
	KC_QUOTE                  // 34   ' and "
	KC_GRAVE                  // 35   Grave accent and tilde
	KC_COMMA                  // 36   , and <
	KC_DOT                    // 37   . and >
	KC_SLASH                  // 38   / and ?
	KC_CAPSLOCK               // 39
	KC_F1                     // 3A
	KC_F2                     // 3B
	KC_F3                     // 3C
	KC_F4                     // 3D
	KC_F5                     // 3E
	KC_F6                     // 3F
	KC_F7                     // 40
	KC_F8                     // 41
	KC_F9                     // 42
	KC_F10                    // 43
	KC_F11                    // 44
	KC_F12                    // 45
	KC_PSCREEN                // 46
	KC_SCROLLLOCK             // 47
	KC_PAUSE                  // 48
	KC_INSERT                 // 49
	KC_HOME                   // 4A
	KC_PGUP                   // 4B
	KC_DELETE                 // 4C
	KC_END                    // 4D
	KC_PGDOWN                 // 4E
	KC_RIGHT                  // 4F
	KC_LEFT                   // 50
	KC_DOWN                   // 51
	KC_UP                     // 52
	KC_NUMLOCK                // 53
	KC_KP_SLASH               // 54
	KC_KP_ASTERISK            // 55
	KC_KP_MINUS               // 56
	KC_KP_PLUS                // 57
	KC_KP_ENTER               // 58
	KC_KP_1                   // 59
	KC_KP_2                   // 5A
	KC_KP_3                   // 5B
	KC_KP_4                   // 5C
	KC_KP_5                   // 5D
	KC_KP_6                   // 5E
	KC_KP_7                   // 5F
	KC_KP_8                   // 60
	KC_KP_9                   // 61
	KC_KP_0                   // 62
	KC_KP_DOT                 // 63
	KC_NONUS_BSLASH           // 64   Non-US \ and | (Typically near the Left-Shift key) */
	KC_APPLICATION            // 65
	KC_POWER                  // 66
	KC_KP_EQUAL               // 67
	KC_F13                    // 68
	KC_F14                    // 69
	KC_F15                    // 6A
	KC_F16                    // 6B
	KC_F17                    // 6C
	KC_F18                    // 6D
	KC_F19                    // 6E
	KC_F20                    // 6F
	KC_F21                    // 70
	KC_F22                    // 71
	KC_F23                    // 72
	KC_F24                    // 73
	KC_EXECUTE                // 74
	KC_HELP                   // 75
	KC_MENU                   // 76
	KC_SELECT                 // 77
	KC_STOP                   // 78
	KC_AGAIN                  // 79
	KC_UNDO                   // 7A
	KC_CUT                    // 7B
	KC_COPY                   // 7C
	KC_PASTE                  // 7D
	KC_FIND                   // 7E
	KC__MUTE                  // 7F
	KC__VOLUP                 // 80
	KC__VOLDOWN               // 81
	KC_LOCKING_CAPS           // 82   locking Caps Lock */
	KC_LOCKING_NUM            // 83   locking Num Lock */
	KC_LOCKING_SCROLL         // 84   locking Scroll Lock */
	KC_KP_COMMA               // 85
	KC_KP_EQUAL_AS400         // 86   equal sign on AS/400 */
	KC_INT1                   // 87
	KC_INT2                   // 88
	KC_INT3                   // 89
	KC_INT4                   // 8A
	KC_INT5                   // 8B
	KC_INT6                   // 8C
	KC_INT7                   // 8D
	KC_INT8                   // 8E
	KC_INT9                   // 8F
	KC_LANG1                  // 90
	KC_LANG2                  // 91
	KC_LANG3                  // 92
	KC_LANG4                  // 93
	KC_LANG5                  // 94
	KC_LANG6                  // 95
	KC_LANG7                  // 96
	KC_LANG8                  // 97
	KC_LANG9                  // 98
	KC_ALT_ERASE              // 99
	KC_SYSREQ                 // 9A
	KC_CANCEL                 // 9B
	KC_CLEAR                  // 9C
	KC_PRIOR                  // 9D
	KC_RETURN                 // 9E
	KC_SEPARATOR              // 9F
	KC_OUT                    // A0
	KC_OPER                   // A1
	KC_CLEAR_AGAIN            // A2
	KC_CRSEL                  // A3
	KC_EXSEL                  // A4
	_                         // A5
	_                         // A6
	_                         // A7
	_                         // A8
	_                         // A9
	_                         // AA
	_                         // AB
	_                         // AC
	_                         // AD
	_                         // AE
	_                         // AF

	/* NOTE: Following code range(0xB0-DD) are shared with special codes of 8-bit keymap */
	KC_KP_00               // B0
	KC_KP_000              // B1
	KC_THOUSANDS_SEPARATOR // B2
	KC_DECIMAL_SEPARATOR   // B3
	KC_CURRENCY_UNIT       // B4
	KC_CURRENCY_SUB_UNIT   // B5
	KC_KP_LPAREN           // B6
	KC_KP_RPAREN           // B7
	KC_KP_LCBRACKET        // B8   {
	KC_KP_RCBRACKET        // B9   }
	KC_KP_TAB              // BA
	KC_KP_BSPACE           // BB
	KC_KP_A                // BC
	KC_KP_B                // BD
	KC_KP_C                // BE
	KC_KP_D                // BF
	KC_KP_E                // C0
	KC_KP_F                // C1
	KC_KP_XOR              // C2
	KC_KP_HAT              // C3
	KC_KP_PERC             // C4
	KC_KP_LT               // C5
	KC_KP_GT               // C6
	KC_KP_AND              // C7
	KC_KP_LAZYAND          // C8
	KC_KP_OR               // C9
	KC_KP_LAZYOR           // CA
	KC_KP_COLON            // CB
	KC_KP_HASH             // CC
	KC_KP_SPACE            // CD
	KC_KP_ATMARK           // CE
	KC_KP_EXCLAMATION      // CF
	KC_KP_MEM_STORE        // D0
	KC_KP_MEM_RECALL       // D1
	KC_KP_MEM_CLEAR        // D2
	KC_KP_MEM_ADD          // D3
	KC_KP_MEM_SUB          // D4
	KC_KP_MEM_MUL          // D5
	KC_KP_MEM_DIV          // D6
	KC_KP_PLUS_MINUS       // D7
	KC_KP_CLEAR            // D8
	KC_KP_CLEAR_ENTRY      // D9
	KC_KP_BINARY           // DA
	KC_KP_OCTAL            // DB
	KC_KP_DECIMAL          // DC
	KC_KP_HEXADECIMAL      // DD
	_                      // DE
	_                      // DF

	/* Modifiers */
	KC_LCTRL  // E0
	KC_LSHIFT // E1
	KC_LALT   // E2
	KC_LGUI   // E3
	KC_RCTRL  // E4
	KC_RSHIFT // E5
	KC_RALT   // E6
	KC_RGUI   // E7
)

var Modifiers = map[string]KeyCode{
	"KC_LCTRL": KC_LCTRL, // E0
	//KC_LSHIFT, // E1
	//KC_LALT,   // E2
	//KC_LGUI ,  // E3
	//KC_RCTRL,  // E4
	//KC_RSHIFT, // E5
	//KC_RALT,   // E6
	//KC_RGUI,   // E7
}

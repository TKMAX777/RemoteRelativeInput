package keymap

type windowsKeyMap map[uint32]WindowsKey

type WindowsKey struct {
	Constant    string
	Value       uint32
	Description string
	EventType   EV_TYPE
	EventInput  string
}

type windowsKeyFromEventInput map[string]WindowsKey

var windowsKeysFromEventInput = windowsKeyFromEventInput{}

func init() {
	for _, k := range windowsKeys {
		windowsKeysFromEventInput[k.EventInput] = k
	}
}

// Windows virtual keycode to key detail
var windowsKeys = windowsKeyMap{
	0x01: WindowsKey{Constant: "VK_LBUTTON", Value: 0x01, Description: "Left mouse button", EventType: EV_TYPE_MOUSE, EventInput: "left"},
	0x02: WindowsKey{Constant: "VK_RBUTTON", Value: 0x02, Description: "Right mouse button", EventType: EV_TYPE_MOUSE, EventInput: "right"},
	0x03: WindowsKey{Constant: "VK_CANCEL", Value: 0x03, Description: "Control-break processing", EventType: EV_TYPE_MOUSE, EventInput: ""},
	0x04: WindowsKey{Constant: "VK_MBUTTON", Value: 0x04, Description: "Middle mouse button (three-button mouse)", EventType: EV_TYPE_MOUSE, EventInput: "middle"},
	0x05: WindowsKey{Constant: "VK_XBUTTON1", Value: 0x05, Description: "X1 mouse button", EventType: EV_TYPE_MOUSE, EventInput: ""},
	0x06: WindowsKey{Constant: "VK_XBUTTON2", Value: 0x06, Description: "X2 mouse button", EventType: EV_TYPE_MOUSE, EventInput: ""},
	0x07: WindowsKey{Constant: "", Value: 0x07, Description: "Undefined", EventType: EV_TYPE_KEY, EventInput: ""},
	0x08: WindowsKey{Constant: "VK_BACK", Value: 0x08, Description: "BACKSPACE key", EventType: EV_TYPE_KEY, EventInput: "BackSpace"},
	0x09: WindowsKey{Constant: "VK_TAB", Value: 0x09, Description: "TAB key", EventType: EV_TYPE_KEY, EventInput: "Tab"},
	0x0C: WindowsKey{Constant: "VK_CLEAR", Value: 0x0C, Description: "CLEAR key", EventType: EV_TYPE_KEY, EventInput: "Clear"},
	0x0D: WindowsKey{Constant: "VK_RETURN", Value: 0x0D, Description: "ENTER key", EventType: EV_TYPE_KEY, EventInput: "Return"},
	0x10: WindowsKey{Constant: "VK_SHIFT", Value: 0x10, Description: "SHIFT key", EventType: EV_TYPE_KEY, EventInput: "shift"},
	0x11: WindowsKey{Constant: "VK_CONTROL", Value: 0x11, Description: "CTRL key", EventType: EV_TYPE_KEY, EventInput: "ctrl"},
	0x12: WindowsKey{Constant: "VK_MENU", Value: 0x12, Description: "ALT key", EventType: EV_TYPE_KEY, EventInput: "alt"},
	0x13: WindowsKey{Constant: "VK_PAUSE", Value: 0x13, Description: "PAUSE key", EventType: EV_TYPE_KEY, EventInput: "Pause"},
	0x14: WindowsKey{Constant: "VK_CAPITAL", Value: 0x14, Description: "CAPS LOCK key", EventType: EV_TYPE_KEY, EventInput: "Caps_Lock"},
	0x15: WindowsKey{Constant: "VK_KANA", Value: 0x15, Description: "IME Kana mode", EventType: EV_TYPE_KEY, EventInput: "kana_switch"},
	0x16: WindowsKey{Constant: "VK_IME_ON", Value: 0x16, Description: "IME On", EventType: EV_TYPE_KEY, EventInput: ""},
	0x17: WindowsKey{Constant: "VK_JUNJA", Value: 0x17, Description: "IME Junja mode", EventType: EV_TYPE_KEY, EventInput: ""},
	0x18: WindowsKey{Constant: "VK_FINAL", Value: 0x18, Description: "IME final mode", EventType: EV_TYPE_KEY, EventInput: ""},
	0x19: WindowsKey{Constant: "VK_KANJI", Value: 0x19, Description: "IME Kanji mode", EventType: EV_TYPE_KEY, EventInput: "kana_switch"},
	0x1A: WindowsKey{Constant: "VK_IME_OFF", Value: 0x1A, Description: "IME Off", EventType: EV_TYPE_KEY, EventInput: ""},
	0x1B: WindowsKey{Constant: "VK_ESCAPE", Value: 0x1B, Description: "ESC key", EventType: EV_TYPE_KEY, EventInput: "Escape"},
	0x1C: WindowsKey{Constant: "VK_CONVERT", Value: 0x1C, Description: "IME convert", EventType: EV_TYPE_KEY, EventInput: ""},
	0x1D: WindowsKey{Constant: "VK_NONCONVERT", Value: 0x1D, Description: "IME nonconvert", EventType: EV_TYPE_KEY, EventInput: ""},
	0x1E: WindowsKey{Constant: "VK_ACCEPT", Value: 0x1E, Description: "IME accept", EventType: EV_TYPE_KEY, EventInput: ""},
	0x1F: WindowsKey{Constant: "VK_MODECHANGE", Value: 0x1F, Description: "IME mode change request", EventType: EV_TYPE_KEY, EventInput: ""},
	0x20: WindowsKey{Constant: "VK_SPACE", Value: 0x20, Description: "SPACEBAR", EventType: EV_TYPE_KEY, EventInput: "space"},
	0x21: WindowsKey{Constant: "VK_PRIOR", Value: 0x21, Description: "PAGE UP key", EventType: EV_TYPE_KEY, EventInput: "Page_Up"},
	0x22: WindowsKey{Constant: "VK_NEXT", Value: 0x22, Description: "PAGE DOWN key", EventType: EV_TYPE_KEY, EventInput: "Page_Down"},
	0x23: WindowsKey{Constant: "VK_END", Value: 0x23, Description: "END key", EventType: EV_TYPE_KEY, EventInput: "End"},
	0x24: WindowsKey{Constant: "VK_HOME", Value: 0x24, Description: "HOME key", EventType: EV_TYPE_KEY, EventInput: "Home"},
	0x25: WindowsKey{Constant: "VK_LEFT", Value: 0x25, Description: "LEFT ARROW key", EventType: EV_TYPE_KEY, EventInput: "Left"},
	0x26: WindowsKey{Constant: "VK_UP", Value: 0x26, Description: "UP ARROW key", EventType: EV_TYPE_KEY, EventInput: "Up"},
	0x27: WindowsKey{Constant: "VK_RIGHT", Value: 0x27, Description: "RIGHT ARROW key", EventType: EV_TYPE_KEY, EventInput: "Right"},
	0x28: WindowsKey{Constant: "VK_DOWN", Value: 0x28, Description: "DOWN ARROW key", EventType: EV_TYPE_KEY, EventInput: "Down"},
	0x29: WindowsKey{Constant: "VK_SELECT", Value: 0x29, Description: "SELECT key", EventType: EV_TYPE_KEY, EventInput: "Select"},
	0x2A: WindowsKey{Constant: "VK_PRINT", Value: 0x2A, Description: "PRINT key", EventType: EV_TYPE_KEY, EventInput: "Print"},
	0x2B: WindowsKey{Constant: "VK_EXECUTE", Value: 0x2B, Description: "EXECUTE key", EventType: EV_TYPE_KEY, EventInput: "Execute"},
	0x2C: WindowsKey{Constant: "VK_SNAPSHOT", Value: 0x2C, Description: "PRINT SCREEN key", EventType: EV_TYPE_KEY, EventInput: "Print"},
	0x2D: WindowsKey{Constant: "VK_INSERT", Value: 0x2D, Description: "INS key", EventType: EV_TYPE_KEY, EventInput: "Insert"},
	0x2E: WindowsKey{Constant: "VK_DELETE", Value: 0x2E, Description: "DEL key", EventType: EV_TYPE_KEY, EventInput: "Delete"},
	0x2F: WindowsKey{Constant: "VK_HELP", Value: 0x2F, Description: "HELP key", EventType: EV_TYPE_KEY, EventInput: "Help"},
	0x30: WindowsKey{Constant: "0 key", Value: 0x30, Description: "0 key", EventType: EV_TYPE_KEY, EventInput: "0"},
	0x31: WindowsKey{Constant: "1 key", Value: 0x31, Description: "1 key", EventType: EV_TYPE_KEY, EventInput: "1"},
	0x32: WindowsKey{Constant: "2 key", Value: 0x32, Description: "2 key", EventType: EV_TYPE_KEY, EventInput: "2"},
	0x33: WindowsKey{Constant: "3 key", Value: 0x33, Description: "3 key", EventType: EV_TYPE_KEY, EventInput: "3"},
	0x34: WindowsKey{Constant: "4 key", Value: 0x34, Description: "4 key", EventType: EV_TYPE_KEY, EventInput: "4"},
	0x35: WindowsKey{Constant: "5 key", Value: 0x35, Description: "5 key", EventType: EV_TYPE_KEY, EventInput: "5"},
	0x36: WindowsKey{Constant: "6 key", Value: 0x36, Description: "6 key", EventType: EV_TYPE_KEY, EventInput: "6"},
	0x37: WindowsKey{Constant: "7 key", Value: 0x37, Description: "7 key", EventType: EV_TYPE_KEY, EventInput: "7"},
	0x38: WindowsKey{Constant: "8 key", Value: 0x38, Description: "8 key", EventType: EV_TYPE_KEY, EventInput: "8"},
	0x39: WindowsKey{Constant: "9 key", Value: 0x39, Description: "9 key", EventType: EV_TYPE_KEY, EventInput: "9"},
	0x41: WindowsKey{Constant: "A key", Value: 0x41, Description: "A key", EventType: EV_TYPE_KEY, EventInput: "a"},
	0x42: WindowsKey{Constant: "B key", Value: 0x42, Description: "B key", EventType: EV_TYPE_KEY, EventInput: "b"},
	0x43: WindowsKey{Constant: "C key", Value: 0x43, Description: "C key", EventType: EV_TYPE_KEY, EventInput: "c"},
	0x44: WindowsKey{Constant: "D key", Value: 0x44, Description: "D key", EventType: EV_TYPE_KEY, EventInput: "d"},
	0x45: WindowsKey{Constant: "E key", Value: 0x45, Description: "E key", EventType: EV_TYPE_KEY, EventInput: "e"},
	0x46: WindowsKey{Constant: "F key", Value: 0x46, Description: "F key", EventType: EV_TYPE_KEY, EventInput: "f"},
	0x47: WindowsKey{Constant: "G key", Value: 0x47, Description: "G key", EventType: EV_TYPE_KEY, EventInput: "g"},
	0x48: WindowsKey{Constant: "H key", Value: 0x48, Description: "H key", EventType: EV_TYPE_KEY, EventInput: "h"},
	0x49: WindowsKey{Constant: "I key", Value: 0x49, Description: "I key", EventType: EV_TYPE_KEY, EventInput: "i"},
	0x4A: WindowsKey{Constant: "J key", Value: 0x4A, Description: "J key", EventType: EV_TYPE_KEY, EventInput: "j"},
	0x4B: WindowsKey{Constant: "K key", Value: 0x4B, Description: "K key", EventType: EV_TYPE_KEY, EventInput: "k"},
	0x4C: WindowsKey{Constant: "L key", Value: 0x4C, Description: "L key", EventType: EV_TYPE_KEY, EventInput: "l"},
	0x4D: WindowsKey{Constant: "M key", Value: 0x4D, Description: "M key", EventType: EV_TYPE_KEY, EventInput: "m"},
	0x4E: WindowsKey{Constant: "N key", Value: 0x4E, Description: "N key", EventType: EV_TYPE_KEY, EventInput: "n"},
	0x4F: WindowsKey{Constant: "O key", Value: 0x4F, Description: "O key", EventType: EV_TYPE_KEY, EventInput: "o"},
	0x50: WindowsKey{Constant: "P key", Value: 0x50, Description: "P key", EventType: EV_TYPE_KEY, EventInput: "p"},
	0x51: WindowsKey{Constant: "Q key", Value: 0x51, Description: "Q key", EventType: EV_TYPE_KEY, EventInput: "q"},
	0x52: WindowsKey{Constant: "R key", Value: 0x52, Description: "R key", EventType: EV_TYPE_KEY, EventInput: "r"},
	0x53: WindowsKey{Constant: "S key", Value: 0x53, Description: "S key", EventType: EV_TYPE_KEY, EventInput: "s"},
	0x54: WindowsKey{Constant: "T key", Value: 0x54, Description: "T key", EventType: EV_TYPE_KEY, EventInput: "t"},
	0x55: WindowsKey{Constant: "U key", Value: 0x55, Description: "U key", EventType: EV_TYPE_KEY, EventInput: "u"},
	0x56: WindowsKey{Constant: "V key", Value: 0x56, Description: "V key", EventType: EV_TYPE_KEY, EventInput: "v"},
	0x57: WindowsKey{Constant: "W key", Value: 0x57, Description: "W key", EventType: EV_TYPE_KEY, EventInput: "w"},
	0x58: WindowsKey{Constant: "X key", Value: 0x58, Description: "X key", EventType: EV_TYPE_KEY, EventInput: "x"},
	0x59: WindowsKey{Constant: "Y key", Value: 0x59, Description: "Y key", EventType: EV_TYPE_KEY, EventInput: "y"},
	0x5A: WindowsKey{Constant: "Z key", Value: 0x5A, Description: "Z key", EventType: EV_TYPE_KEY, EventInput: "z"},
	0x5B: WindowsKey{Constant: "VK_LWIN", Value: 0x5B, Description: "Left Windows key (Natural keyboard)", EventType: EV_TYPE_KEY, EventInput: "Super_L"},
	0x5C: WindowsKey{Constant: "VK_RWIN", Value: 0x5C, Description: "Right Windows key (Natural keyboard)", EventType: EV_TYPE_KEY, EventInput: "Super_R"},
	0x5D: WindowsKey{Constant: "VK_APPS", Value: 0x5D, Description: "Applications key (Natural keyboard)", EventType: EV_TYPE_KEY, EventInput: ""},
	0x5E: WindowsKey{Constant: "", Value: 0x5E, Description: "Reserved", EventType: EV_TYPE_KEY, EventInput: ""},
	0x5F: WindowsKey{Constant: "VK_SLEEP", Value: 0x5F, Description: "Computer Sleep key", EventType: EV_TYPE_KEY, EventInput: ""},
	0x60: WindowsKey{Constant: "VK_NUMPAD0", Value: 0x60, Description: "Numeric keypad 0 key", EventType: EV_TYPE_KEY, EventInput: "0"},
	0x61: WindowsKey{Constant: "VK_NUMPAD1", Value: 0x61, Description: "Numeric keypad 1 key", EventType: EV_TYPE_KEY, EventInput: "1"},
	0x62: WindowsKey{Constant: "VK_NUMPAD2", Value: 0x62, Description: "Numeric keypad 2 key", EventType: EV_TYPE_KEY, EventInput: "2"},
	0x63: WindowsKey{Constant: "VK_NUMPAD3", Value: 0x63, Description: "Numeric keypad 3 key", EventType: EV_TYPE_KEY, EventInput: "3"},
	0x64: WindowsKey{Constant: "VK_NUMPAD4", Value: 0x64, Description: "Numeric keypad 4 key", EventType: EV_TYPE_KEY, EventInput: "4"},
	0x65: WindowsKey{Constant: "VK_NUMPAD5", Value: 0x65, Description: "Numeric keypad 5 key", EventType: EV_TYPE_KEY, EventInput: "5"},
	0x66: WindowsKey{Constant: "VK_NUMPAD6", Value: 0x66, Description: "Numeric keypad 6 key", EventType: EV_TYPE_KEY, EventInput: "6"},
	0x67: WindowsKey{Constant: "VK_NUMPAD7", Value: 0x67, Description: "Numeric keypad 7 key", EventType: EV_TYPE_KEY, EventInput: "7"},
	0x68: WindowsKey{Constant: "VK_NUMPAD8", Value: 0x68, Description: "Numeric keypad 8 key", EventType: EV_TYPE_KEY, EventInput: "8"},
	0x69: WindowsKey{Constant: "VK_NUMPAD9", Value: 0x69, Description: "Numeric keypad 9 key", EventType: EV_TYPE_KEY, EventInput: "9"},
	0x6A: WindowsKey{Constant: "VK_MULTIPLY", Value: 0x6A, Description: "Multiply key", EventType: EV_TYPE_KEY, EventInput: "multiply"},
	0x6B: WindowsKey{Constant: "VK_ADD", Value: 0x6B, Description: "Add key", EventType: EV_TYPE_KEY, EventInput: ""},
	0x6C: WindowsKey{Constant: "VK_SEPARATOR", Value: 0x6C, Description: "Separator key", EventType: EV_TYPE_KEY, EventInput: "bar"},
	0x6D: WindowsKey{Constant: "VK_SUBTRACT", Value: 0x6D, Description: "Subtract key", EventType: EV_TYPE_KEY, EventInput: "minus"},
	0x6E: WindowsKey{Constant: "VK_DECIMAL", Value: 0x6E, Description: "Decimal key", EventType: EV_TYPE_KEY, EventInput: "period"},
	0x6F: WindowsKey{Constant: "VK_DIVIDE", Value: 0x6F, Description: "Divide key", EventType: EV_TYPE_KEY, EventInput: "slash"},
	0x70: WindowsKey{Constant: "VK_F1", Value: 0x70, Description: "F1 key", EventType: EV_TYPE_KEY, EventInput: "F1"},
	0x71: WindowsKey{Constant: "VK_F2", Value: 0x71, Description: "F2 key", EventType: EV_TYPE_KEY, EventInput: "F2"},
	0x72: WindowsKey{Constant: "VK_F3", Value: 0x72, Description: "F3 key", EventType: EV_TYPE_KEY, EventInput: "F3"},
	0x73: WindowsKey{Constant: "VK_F4", Value: 0x73, Description: "F4 key", EventType: EV_TYPE_KEY, EventInput: "F4"},
	0x74: WindowsKey{Constant: "VK_F5", Value: 0x74, Description: "F5 key", EventType: EV_TYPE_KEY, EventInput: "F5"},
	0x75: WindowsKey{Constant: "VK_F6", Value: 0x75, Description: "F6 key", EventType: EV_TYPE_KEY, EventInput: "F6"},
	0x76: WindowsKey{Constant: "VK_F7", Value: 0x76, Description: "F7 key", EventType: EV_TYPE_KEY, EventInput: "F7"},
	0x77: WindowsKey{Constant: "VK_F8", Value: 0x77, Description: "F8 key", EventType: EV_TYPE_KEY, EventInput: "F8"},
	0x78: WindowsKey{Constant: "VK_F9", Value: 0x78, Description: "F9 key", EventType: EV_TYPE_KEY, EventInput: "F9"},
	0x79: WindowsKey{Constant: "VK_F10", Value: 0x79, Description: "F10 key", EventType: EV_TYPE_KEY, EventInput: "F10"},
	0x7A: WindowsKey{Constant: "VK_F11", Value: 0x7A, Description: "F11 key", EventType: EV_TYPE_KEY, EventInput: "F11"},
	0x7B: WindowsKey{Constant: "VK_F12", Value: 0x7B, Description: "F12 key", EventType: EV_TYPE_KEY, EventInput: "F12"},
	0x7C: WindowsKey{Constant: "VK_F13", Value: 0x7C, Description: "F13 key", EventType: EV_TYPE_KEY, EventInput: "F13"},
	0x7D: WindowsKey{Constant: "VK_F14", Value: 0x7D, Description: "F14 key", EventType: EV_TYPE_KEY, EventInput: "F14"},
	0x7E: WindowsKey{Constant: "VK_F15", Value: 0x7E, Description: "F15 key", EventType: EV_TYPE_KEY, EventInput: "F15"},
	0x7F: WindowsKey{Constant: "VK_F16", Value: 0x7F, Description: "F16 key", EventType: EV_TYPE_KEY, EventInput: "F16"},
	0x80: WindowsKey{Constant: "VK_F17", Value: 0x80, Description: "F17 key", EventType: EV_TYPE_KEY, EventInput: "F17"},
	0x81: WindowsKey{Constant: "VK_F18", Value: 0x81, Description: "F18 key", EventType: EV_TYPE_KEY, EventInput: "F18"},
	0x82: WindowsKey{Constant: "VK_F19", Value: 0x82, Description: "F19 key", EventType: EV_TYPE_KEY, EventInput: "F19"},
	0x83: WindowsKey{Constant: "VK_F20", Value: 0x83, Description: "F20 key", EventType: EV_TYPE_KEY, EventInput: "F20"},
	0x84: WindowsKey{Constant: "VK_F21", Value: 0x84, Description: "F21 key", EventType: EV_TYPE_KEY, EventInput: "F21"},
	0x85: WindowsKey{Constant: "VK_F22", Value: 0x85, Description: "F22 key", EventType: EV_TYPE_KEY, EventInput: "F22"},
	0x86: WindowsKey{Constant: "VK_F23", Value: 0x86, Description: "F23 key", EventType: EV_TYPE_KEY, EventInput: "F23"},
	0x87: WindowsKey{Constant: "VK_F24", Value: 0x87, Description: "F24 key", EventType: EV_TYPE_KEY, EventInput: "F24"},
	0x90: WindowsKey{Constant: "VK_NUMLOCK", Value: 0x90, Description: "NUM LOCK key", EventType: EV_TYPE_KEY, EventInput: "Num_Lock"},
	0x91: WindowsKey{Constant: "VK_SCROLL", Value: 0x91, Description: "SCROLL LOCK key", EventType: EV_TYPE_KEY, EventInput: "Scroll_Lock"},
	0xA0: WindowsKey{Constant: "VK_LSHIFT", Value: 0xA0, Description: "Left SHIFT key", EventType: EV_TYPE_KEY, EventInput: "shift"},
	0xA1: WindowsKey{Constant: "VK_RSHIFT", Value: 0xA1, Description: "Right SHIFT key", EventType: EV_TYPE_KEY, EventInput: "shift"},
	0xA2: WindowsKey{Constant: "VK_LCONTROL", Value: 0xA2, Description: "Left CONTROL key", EventType: EV_TYPE_KEY, EventInput: "ctrl"},
	0xA3: WindowsKey{Constant: "VK_RCONTROL", Value: 0xA3, Description: "Right CONTROL key", EventType: EV_TYPE_KEY, EventInput: "ctrl"},
	0xA4: WindowsKey{Constant: "VK_LMENU", Value: 0xA4, Description: "Left MENU key", EventType: EV_TYPE_KEY, EventInput: "Menu"},
	0xA5: WindowsKey{Constant: "VK_RMENU", Value: 0xA5, Description: "Right MENU key", EventType: EV_TYPE_KEY, EventInput: "Menu"},
	0xA6: WindowsKey{Constant: "VK_BROWSER_BACK", Value: 0xA6, Description: "Browser Back key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xA7: WindowsKey{Constant: "VK_BROWSER_FORWARD", Value: 0xA7, Description: "Browser Forward key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xA8: WindowsKey{Constant: "VK_BROWSER_REFRESH", Value: 0xA8, Description: "Browser Refresh key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xA9: WindowsKey{Constant: "VK_BROWSER_STOP", Value: 0xA9, Description: "Browser Stop key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xAA: WindowsKey{Constant: "VK_BROWSER_SEARCH", Value: 0xAA, Description: "Browser Search key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xAB: WindowsKey{Constant: "VK_BROWSER_FAVORITES", Value: 0xAB, Description: "Browser Favorites key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xAC: WindowsKey{Constant: "VK_BROWSER_HOME", Value: 0xAC, Description: "Browser Start and Home key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xAD: WindowsKey{Constant: "VK_VOLUME_MUTE", Value: 0xAD, Description: "Volume Mute key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xAE: WindowsKey{Constant: "VK_VOLUME_DOWN", Value: 0xAE, Description: "Volume Down key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xAF: WindowsKey{Constant: "VK_VOLUME_UP", Value: 0xAF, Description: "Volume Up key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xB0: WindowsKey{Constant: "VK_MEDIA_NEXT_TRACK", Value: 0xB0, Description: "Next Track key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xB1: WindowsKey{Constant: "VK_MEDIA_PREV_TRACK", Value: 0xB1, Description: "Previous Track key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xB2: WindowsKey{Constant: "VK_MEDIA_STOP", Value: 0xB2, Description: "Stop Media key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xB3: WindowsKey{Constant: "VK_MEDIA_PLAY_PAUSE", Value: 0xB3, Description: "Play/Pause Media key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xB4: WindowsKey{Constant: "VK_LAUNCH_MAIL", Value: 0xB4, Description: "Start Mail key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xB5: WindowsKey{Constant: "VK_LAUNCH_MEDIA_SELECT", Value: 0xB5, Description: "Select Media key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xB6: WindowsKey{Constant: "VK_LAUNCH_APP1", Value: 0xB6, Description: "Start Application 1 key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xB7: WindowsKey{Constant: "VK_LAUNCH_APP2", Value: 0xB7, Description: "Start Application 2 key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xBA: WindowsKey{Constant: "VK_OEM_1", Value: 0xBA, Description: "Used for miscellaneous characters; it can vary by keyboard.", EventType: EV_TYPE_KEY, EventInput: ""},
	0xBB: WindowsKey{Constant: "VK_OEM_PLUS", Value: 0xBB, Description: "For any country/region, the '+' key", EventType: EV_TYPE_KEY, EventInput: "semicolon"},
	0xBC: WindowsKey{Constant: "VK_OEM_COMMA", Value: 0xBC, Description: "For any country/region, the ',' key", EventType: EV_TYPE_KEY, EventInput: "comma"},
	0xBD: WindowsKey{Constant: "VK_OEM_MINUS", Value: 0xBD, Description: "For any country/region, the '-' key", EventType: EV_TYPE_KEY, EventInput: "minus"},
	0xBE: WindowsKey{Constant: "VK_OEM_PERIOD", Value: 0xBE, Description: "For any country/region, the '.' key", EventType: EV_TYPE_KEY, EventInput: "period"},
	0xBF: WindowsKey{Constant: "VK_OEM_2", Value: 0xBF, Description: "Used for miscellaneous characters; it can vary by keyboard.", EventType: EV_TYPE_KEY, EventInput: ""},
	0xC0: WindowsKey{Constant: "VK_OEM_3", Value: 0xC0, Description: "Used for miscellaneous characters; it can vary by keyboard.", EventType: EV_TYPE_KEY, EventInput: ""},
	0xDB: WindowsKey{Constant: "VK_OEM_4", Value: 0xDB, Description: "Used for miscellaneous characters; it can vary by keyboard.", EventType: EV_TYPE_KEY, EventInput: "parenleft"},
	0xDC: WindowsKey{Constant: "VK_OEM_5", Value: 0xDC, Description: "Used for miscellaneous characters; it can vary by keyboard.", EventType: EV_TYPE_KEY, EventInput: "backslash"},
	0xDD: WindowsKey{Constant: "VK_OEM_6", Value: 0xDD, Description: "Used for miscellaneous characters; it can vary by keyboard.", EventType: EV_TYPE_KEY, EventInput: "parenright"},
	0xDE: WindowsKey{Constant: "VK_OEM_7", Value: 0xDE, Description: "Used for miscellaneous characters; it can vary by keyboard.", EventType: EV_TYPE_KEY, EventInput: "caret"},
	0xDF: WindowsKey{Constant: "VK_OEM_8", Value: 0xDF, Description: "Used for miscellaneous characters; it can vary by keyboard.", EventType: EV_TYPE_KEY, EventInput: ""},
	0xE0: WindowsKey{Constant: "", Value: 0xE0, Description: "Reserved", EventType: EV_TYPE_KEY, EventInput: ""},
	0xE1: WindowsKey{Constant: "", Value: 0xE1, Description: "OEM specific", EventType: EV_TYPE_KEY, EventInput: ""},
	0xE2: WindowsKey{Constant: "VK_OEM_102", Value: 0xE2, Description: "Either the angle bracket key or the backslash key on the RT 102-key keyboard", EventType: EV_TYPE_KEY, EventInput: "underscore"},
	0xE5: WindowsKey{Constant: "VK_PROCESSKEY", Value: 0xE5, Description: "IME PROCESS key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xE6: WindowsKey{Constant: "", Value: 0xE6, Description: "OEM specific", EventType: EV_TYPE_KEY, EventInput: ""},
	0xE7: WindowsKey{Constant: "VK_PACKET", Value: 0xE7, Description: "Used to pass Unicode characters as if they were keystrokes. The VK_PACKET key is the low word of a 32-bit Virtual Key value used for non-keyboard input methods. For more information, see Remark in KEYBDINPUT, SendInput, WM_KEYDOWN, and WM_KEYUP", EventType: EV_TYPE_KEY, EventInput: ""},
	0xE8: WindowsKey{Constant: "", Value: 0xE8, Description: "Unassigned", EventType: EV_TYPE_KEY, EventInput: ""},
	0xF6: WindowsKey{Constant: "VK_ATTN", Value: 0xF6, Description: "Attn key", EventType: EV_TYPE_KEY, EventInput: "3270_Attn"},
	0xF7: WindowsKey{Constant: "VK_CRSEL", Value: 0xF7, Description: "CrSel key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xF8: WindowsKey{Constant: "VK_EXSEL", Value: 0xF8, Description: "ExSel key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xF9: WindowsKey{Constant: "VK_EREOF", Value: 0xF9, Description: "Erase EOF key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xFA: WindowsKey{Constant: "VK_PLAY", Value: 0xFA, Description: "Play key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xFB: WindowsKey{Constant: "VK_ZOOM", Value: 0xFB, Description: "Zoom key", EventType: EV_TYPE_KEY, EventInput: ""},
	0xFC: WindowsKey{Constant: "VK_NONAME", Value: 0xFC, Description: "Reserved", EventType: EV_TYPE_KEY, EventInput: ""},
	0xFD: WindowsKey{Constant: "VK_PA1", Value: 0xFD, Description: "PA1 key", EventType: EV_TYPE_KEY, EventInput: "PA1"},
	0xFE: WindowsKey{Constant: "VK_OEM_CLEAR", Value: 0xFE, Description: "Clear key", EventType: EV_TYPE_KEY, EventInput: "Clear"},
}
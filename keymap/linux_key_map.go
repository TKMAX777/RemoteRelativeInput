package keymap

type linuxKeyMap map[uint32]LinuxKey

type LinuxKey struct {
	Constant    string
	Value       uint32
	Description string
	WindowsKey  uint32
	EventInput  string
}

type linuxKeyMapFromEventInput map[string]LinuxKey

var linuxKeysFromEventInput = linuxKeyMapFromEventInput{}

func init() {
	for _, k := range linuxKeys {
		linuxKeysFromEventInput[k.EventInput] = k
	}
}

// Linux keycode to key details
var linuxKeys = linuxKeyMap{
	53:  {Constant: "Esc", Value: 53, Description: "Esc", EventInput: "Escape"},
	122: {Constant: "F1", Value: 122, Description: "F1", EventInput: "F1"},
	120: {Constant: "F2", Value: 120, Description: "F2", EventInput: "F2"},
	99:  {Constant: "F3", Value: 99, Description: "F3", EventInput: "F3"},
	118: {Constant: "F4", Value: 118, Description: "F4", EventInput: "F4"},
	96:  {Constant: "F5", Value: 96, Description: "F5", EventInput: "F5"},
	97:  {Constant: "F6", Value: 97, Description: "F6", EventInput: "F6"},
	98:  {Constant: "F7", Value: 98, Description: "F7", EventInput: "F7"},
	100: {Constant: "F8", Value: 100, Description: "F8", EventInput: "F8"},
	101: {Constant: "F9", Value: 101, Description: "F9", EventInput: "F9"},
	109: {Constant: "F10", Value: 109, Description: "F10", EventInput: "F10"},
	103: {Constant: "F11", Value: 103, Description: "F11", EventInput: "F11"},
	127: {Constant: "F12", Value: 127, Description: "F12", EventInput: "F12"},
	105: {Constant: "PrintScrn", Value: 105, Description: "PrintScrn", EventInput: "Print"},
	107: {Constant: "Scroll Lock", Value: 107, Description: "Scroll Lock", EventInput: "Scroll_Lock"},
	113: {Constant: "Pause", Value: 113, Description: "Pause", EventInput: "Pause"},
	50:  {Constant: "`", Value: 50, Description: "`"},
	18:  {Constant: "1", Value: 18, Description: "1", EventInput: "1"},
	19:  {Constant: "2", Value: 19, Description: "2", EventInput: "2"},
	20:  {Constant: "3", Value: 20, Description: "3", EventInput: "3"},
	21:  {Constant: "4", Value: 21, Description: "4", EventInput: "4"},
	23:  {Constant: "5", Value: 23, Description: "5", EventInput: "5"},
	22:  {Constant: "6", Value: 22, Description: "6", EventInput: "6"},
	26:  {Constant: "7", Value: 26, Description: "7", EventInput: "7"},
	28:  {Constant: "8", Value: 28, Description: "8", EventInput: "8"},
	25:  {Constant: "9", Value: 25, Description: "9", EventInput: "9"},
	29:  {Constant: "0", Value: 29, Description: "0", EventInput: "0"},
	27:  {Constant: "-", Value: 27, Description: "-", EventInput: "minus"},
	24:  {Constant: "=", Value: 24, Description: "="},
	51:  {Constant: "Backspace", Value: 51, Description: "Backspace", EventInput: "BackSpace"},
	114: {Constant: "Insert", Value: 114, Description: "Insert", EventInput: "Insert"},
	115: {Constant: "Home", Value: 115, Description: "Home", EventInput: "Home"},
	116: {Constant: "Page Up", Value: 116, Description: "Page Up", EventInput: "Page_Up"},
	71:  {Constant: "Num Lock", Value: 71, Description: "Num Lock", EventInput: "Num_Lock"},
	75:  {Constant: "KP /", Value: 75, Description: "KP /"},
	67:  {Constant: "KP *", Value: 67, Description: "KP *"},
	78:  {Constant: "KP -", Value: 78, Description: "KP -"},
	48:  {Constant: "Tab", Value: 48, Description: "Tab", EventInput: "Tab"},
	12:  {Constant: "Q", Value: 12, Description: "Q", EventInput: "q"},
	13:  {Constant: "W", Value: 13, Description: "W", EventInput: "w"},
	14:  {Constant: "E", Value: 14, Description: "E", EventInput: "e"},
	15:  {Constant: "R", Value: 15, Description: "R", EventInput: "r"},
	17:  {Constant: "T", Value: 17, Description: "T", EventInput: "t"},
	16:  {Constant: "Y", Value: 16, Description: "Y", EventInput: "y"},
	32:  {Constant: "U", Value: 32, Description: "U", EventInput: "u"},
	34:  {Constant: "I", Value: 34, Description: "I", EventInput: "i"},
	31:  {Constant: "O", Value: 31, Description: "O", EventInput: "o"},
	35:  {Constant: "P", Value: 35, Description: "P", EventInput: "p"},
	33:  {Constant: "[", Value: 33, Description: "[", EventInput: "["},
	30:  {Constant: "]", Value: 30, Description: "]", EventInput: "]"},
	36:  {Constant: "Return", Value: 36, Description: "Return", EventInput: "Return"},
	117: {Constant: "Delete", Value: 117, Description: "Delete", EventInput: "Delete"},
	119: {Constant: "End", Value: 119, Description: "End", EventInput: "End"},
	121: {Constant: "Page Down", Value: 121, Description: "Page Down", EventInput: "Page_Down"},
	89:  {Constant: "KP 7", Value: 89, Description: "KP 7", EventInput: ""},
	91:  {Constant: "KP 8", Value: 91, Description: "KP 8", EventInput: ""},
	92:  {Constant: "KP 9", Value: 92, Description: "KP 9", EventInput: ""},
	69:  {Constant: "KP +", Value: 69, Description: "KP +", EventInput: ""},
	57:  {Constant: "Caps Lock", Value: 57, Description: "Caps Lock", EventInput: "Caps_Lock"},
	0:   {Constant: "A", Value: 0, Description: "A", EventInput: "a"},
	1:   {Constant: "S", Value: 1, Description: "S", EventInput: "s"},
	2:   {Constant: "D", Value: 2, Description: "D", EventInput: "d"},
	3:   {Constant: "F", Value: 3, Description: "F", EventInput: "f"},
	5:   {Constant: "G", Value: 5, Description: "G", EventInput: "g"},
	4:   {Constant: "H", Value: 4, Description: "H", EventInput: "h"},
	38:  {Constant: "J", Value: 38, Description: "J", EventInput: "j"},
	40:  {Constant: "K", Value: 40, Description: "K", EventInput: "k"},
	37:  {Constant: "L", Value: 37, Description: "L", EventInput: "l"},
	41:  {Constant: ";", Value: 41, Description: ";", EventInput: ";"},
	39:  {Constant: "'", Value: 39, Description: "'", EventInput: "'"},
	86:  {Constant: "KP 4", Value: 86, Description: "KP 4", EventInput: ""},
	87:  {Constant: "KP 5", Value: 87, Description: "KP 5", EventInput: ""},
	88:  {Constant: "KP 6", Value: 88, Description: "KP 6", EventInput: ""},
	56:  {Constant: "Shift Left", Value: 56, Description: "Shift Left", EventInput: ""},
	6:   {Constant: "Z", Value: 6, Description: "Z", EventInput: "z"},
	7:   {Constant: "X", Value: 7, Description: "X", EventInput: "x"},
	8:   {Constant: "C", Value: 8, Description: "C", EventInput: "c"},
	9:   {Constant: "V", Value: 9, Description: "V", EventInput: "v"},
	11:  {Constant: "B", Value: 11, Description: "B", EventInput: "b"},
	45:  {Constant: "N", Value: 45, Description: "N", EventInput: "n"},
	46:  {Constant: "M", Value: 46, Description: "M", EventInput: "m"},
	43:  {Constant: ",", Value: 43, Description: ",", EventInput: "comma"},
	47:  {Constant: ".", Value: 47, Description: ".", EventInput: "period"},
	44:  {Constant: "/", Value: 44, Description: "/", EventInput: "slash"},
	42:  {Constant: "\\", Value: 42, Description: "\\", EventInput: "backslash"},
	62:  {Constant: "Cursor Up", Value: 62, Description: "Cursor Up", EventInput: ""},
	83:  {Constant: "KP 1", Value: 83, Description: "KP 1", EventInput: ""},
	84:  {Constant: "KP 2", Value: 84, Description: "KP 2", EventInput: ""},
	85:  {Constant: "KP 3", Value: 85, Description: "KP 3", EventInput: ""},
	76:  {Constant: "KP Enter", Value: 76, Description: "KP Enter", EventInput: ""},
	54:  {Constant: "Ctrl Left", Value: 54, Description: "Ctrl Left", EventInput: ""},
	58:  {Constant: "Logo Left (-> Option)", Value: 58, Description: "Logo Left (-> Option)", EventInput: ""},
	55:  {Constant: "Alt Left (-> Command)", Value: 55, Description: "Alt Left (-> Command)", EventInput: ""},
	49:  {Constant: "Space", Value: 49, Description: "Space", EventInput: ""},
	59:  {Constant: "Cursor Left", Value: 59, Description: "Cursor Left", EventInput: "ctrl"},
	61:  {Constant: "Cursor Down", Value: 61, Description: "Cursor Down", EventInput: "ctrl"},
	60:  {Constant: "Cursor Right", Value: 60, Description: "Cursor Right", EventInput: ""},
	82:  {Constant: "KP 0", Value: 82, Description: "KP 0", EventInput: ""},
	65:  {Constant: "KP .", Value: 65, Description: "KP .", EventInput: ""},
}

package telex

type ToneMark rune

const (
	Acute     ToneMark = 's'
	Grave     ToneMark = 'f'
	HookAbove ToneMark = 'r'
	Tilde     ToneMark = 'x'
	Underdot  ToneMark = 'j'

	None ToneMark = ' '
)

func runeToToneMark(ch rune) ToneMark {
	switch ch {
	case 's':
		return Acute
	case 'f':
		return Grave
	case 'r':
		return HookAbove
	case 'x':
		return Tilde
	case 'j':
		return Underdot
	default:
		return None
	}
}

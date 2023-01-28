package telex

import (
	"strings"

	"github.com/thanhfphan/vigo/utils"
)

func addTone(word string, t ToneMark) string {
	var result strings.Builder
	clean := utils.RemoveToneMark(word)
	if tone := findToneMark(word); tone == t {
		result.WriteString(clean)
		result.WriteRune(rune(tone))
		return result.String()
	}

	markIndex, err := utils.GetToneMarkPlacement(word)
	if err != nil {
		return word
	}

	for i, ch := range clean {
		if i != markIndex {
			result.WriteRune(ch)
			continue
		}

		m := getToneMarkMap(t)
		toneMarkCh := clean[markIndex]
		if v, ok := m[rune(toneMarkCh)]; ok {
			result.WriteRune(v)
		} else {
			result.WriteRune(ch)
		}
	}

	return result.String()
}

func getToneMarkMap(t ToneMark) map[rune]rune {
	switch t {
	case Acute:
		return utils.AccuteMap
	case Grave:
		return utils.GraveMap
	case HookAbove:
		return utils.HookAboveMap
	case Tilde:
		return utils.TildeMap
	case Underdot:
		return utils.DotMap
	default:
		return map[rune]rune{}
	}
}

func findToneMarkRune(r rune) ToneMark {

	for _, v := range utils.AccuteMap {
		if v == r {
			return Acute
		}
	}

	for _, v := range utils.GraveMap {
		if v == r {
			return Grave
		}
	}

	for _, v := range utils.HookAboveMap {
		if v == r {
			return HookAbove
		}
	}

	for _, v := range utils.TildeMap {
		if v == r {
			return Tilde
		}
	}

	for _, v := range utils.DotMap {
		if v == r {
			return Underdot
		}
	}

	return None
}

func findToneMark(input string) ToneMark {
	for _, ch := range input {
		tm := findToneMarkRune(ch)
		if tm != None {
			return tm
		}
	}

	return None
}

func transformWord(word string) string {
	var result strings.Builder
	for _, ch := range word {
		switch ch {
		case 's', 'f', 'r', 'x', 'j':
			tmp := addTone(result.String(), runeToToneMark(ch))
			result.Reset()
			result.WriteString(tmp)
		// case 'a':
		// case 'e':
		// case 'o':
		// 	// TODO

		// case 'w':
		// case 'd':
		// TODO
		default:
			result.WriteRune(ch)
		}
	}

	return result.String()
}

package utils

import (
	"errors"
	"strings"
)

func RemoveToneMarkRune(ch rune) rune {
	for k, v := range ToneMarkMap {
		if strings.ContainsRune(v, ch) {
			return k
		}
	}

	return ch
}

func IsVowels(r rune) bool {
	switch r {
	case 'a', 'ă', 'â', 'e', 'ê', 'i', 'o', 'ô', 'ơ', 'u', 'ư', 'y':
		return true
	default:
		return false
	}
}

func IsModifiedVowels(r rune) bool {
	switch r {
	case 'ă', 'â', 'ê', 'ô', 'ơ', 'ư':
		return true
	default:
		return false
	}
}

func IsModifiableVowels(r rune) bool {
	switch r {
	case 'a', 'e', 'o', 'u':
		return true
	default:
		return false
	}
}

func RemoveToneMark(word string) string {
	var result strings.Builder

	for _, ch := range word {
		result.WriteRune(RemoveToneMarkRune(ch))
	}

	return result.String()
}

// returns [start] and [end] indx of vowels
// returns error if not found
func GetIndexVowels(word string) (int, int, error) {
	word = strings.ToLower(word)
	start, end := -1, -1

	for i, ch := range word {
		if IsVowels(ch) {
			start = i
			end = i
			for _, ch := range word[i:] {
				if IsVowels(ch) {
					end++
				}
			}
			break
		}
	}
	if start == -1 || end == -1 {
		return -1, -1, errors.New("not found vowels")
	}

	if word[0] == 'q' && start == 1 && word[start] == 'u' {
		start++
	} else if word[0] == 'g' && start == 1 && word[start] == 'i' {
		start++
	}

	if start > end {
		return -1, -1, errors.New("not found vowels")
	}

	return start, end, nil
}

func GetToneMarkPlacement(word string) (int, error) {
	lenghtW := len(word)
	sIdx, eIdx, err := GetIndexVowels(word)
	if err != nil {
		return -1, err
	}

	if sIdx == eIdx {
		return sIdx, nil
	}

	vowels := word[sIdx:eIdx]
	for i, ch := range vowels {
		if ch == 'ơ' {
			return sIdx + i, nil
		}
	}

	for i, ch := range vowels {
		if IsModifiedVowels(ch) {
			return sIdx + i, nil
		}
	}

	if word[sIdx] == 'o' && sIdx+1 < lenghtW {
		switch word[sIdx+1] {
		case 'a', 'e', 'o':
			return sIdx + 1, nil
		}
	}

	if word[sIdx] == 'u' && sIdx+1 < lenghtW && word[sIdx+1] == 'y' {
		return sIdx + 1, nil
	}

	if eIdx == lenghtW-1 && eIdx > sIdx {
		return eIdx - 1, nil
	}

	for i, ch := range vowels {
		if IsModifiableVowels(ch) {
			return sIdx + i, nil
		}
	}

	return sIdx, nil
}

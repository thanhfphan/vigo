package vigo

import (
	"github.com/thanhfphan/vigo/telex"
)

func TelexStyle(input string) string {
	return telex.Transform(input)
}

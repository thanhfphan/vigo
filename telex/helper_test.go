package telex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addTone(t *testing.T) {
	r := require.New(t)

	tcs := []struct {
		Input  string
		TM     ToneMark
		Expect string
	}{
		{
			Input:  "qua",
			TM:     Acute,
			Expect: "quá",
		},
		{
			Input:  "banh",
			TM:     Acute,
			Expect: "bánh",
		},
		{
			Input:  "banh",
			TM:     Grave,
			Expect: "bành",
		},
		{
			Input:  "nga",
			TM:     HookAbove,
			Expect: "ngả",
		},
		{
			Input:  "nga",
			TM:     Tilde,
			Expect: "ngã",
		},
		{
			Input:  "dam",
			TM:     Underdot,
			Expect: "dạm",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Input, func(t *testing.T) {
			result := addTone(tc.Input, tc.TM)
			r.Equal(tc.Expect, result)
		})
	}
}

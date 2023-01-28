package telex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Tranform(t *testing.T) {
	r := require.New(t)

	tcs := []struct {
		Input  string
		Expect string
	}{
		{
			Input:  "qua",
			Expect: "qua",
		},
		{
			Input:  "quas",
			Expect: "quá",
		},
		{
			Input:  "quas nef",
			Expect: "quá nè",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Input, func(t *testing.T) {
			result := Transform(tc.Input)
			r.Equal(tc.Expect, result)
		})
	}
}

package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RemoveToneMark(t *testing.T) {
	r := require.New(t)
	tcs := []struct {
		Input  string
		Expect string
	}{
		{
			Input:  "quả",
			Expect: "qua",
		},
		{
			Input:  "bưởi",
			Expect: "bươi",
		},
		{
			Input:  "yểu",
			Expect: "yêu",
		},
		{
			Input:  "điệu",
			Expect: "điêu",
		},
		{
			Input:  "giận",
			Expect: "giân",
		},
		{
			Input:  "dỗi",
			Expect: "dôi",
		},
		{
			Input:  "bánh",
			Expect: "banh",
		},
		{
			Input:  "cuốn",
			Expect: "cuôn",
		},
		{
			Input:  "nước",
			Expect: "nươc",
		},
		{
			Input:  "mắm",
			Expect: "măm",
		},
		{
			Input:  "tuyệt,",
			Expect: "tuyêt,",
		},
		{
			Input:  "tuyệt!",
			Expect: "tuyêt!",
		},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%s", tc.Input), func(t *testing.T) {
			result := RemoveToneMark(tc.Input)
			r.Equal(tc.Expect, result)
		})
	}
}

func Test_GetToneMarkPlacement(t *testing.T) {
	r := require.New(t)
	tcs := []struct {
		Input          string
		ExpectIndex    int
		ExpectHasError bool
	}{
		{
			Input:          "quan",
			ExpectIndex:    2,
			ExpectHasError: false,
		},
		{
			Input:          "choe",
			ExpectIndex:    3,
			ExpectHasError: false,
		},
		{
			Input:          "thanh",
			ExpectIndex:    2,
			ExpectHasError: false,
		},
		{
			Input:          "tttt",
			ExpectIndex:    -1,
			ExpectHasError: true,
		},
		{
			Input:          "gia",
			ExpectIndex:    2,
			ExpectHasError: false,
		},
		{
			Input:          "giau",
			ExpectIndex:    2,
			ExpectHasError: false,
		},
		{
			Input:          "cau",
			ExpectIndex:    1,
			ExpectHasError: false,
		},
		{
			Input:          "quy",
			ExpectIndex:    2,
			ExpectHasError: false,
		},
		{
			Input:          "choa",
			ExpectIndex:    3,
			ExpectHasError: false,
		},
		{
			Input:          "dam",
			ExpectIndex:    1,
			ExpectHasError: false,
		},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("Test the word [%s]", tc.Input), func(t *testing.T) {
			result, err := GetToneMarkPlacement(tc.Input)
			if tc.ExpectHasError {
				r.Error(err)
			} else {
				r.NoError(err)
			}
			r.Equal(tc.ExpectIndex, result)
		})
	}
}

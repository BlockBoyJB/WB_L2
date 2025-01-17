package main

import (
	"errors"
	"testing"
)

func TestUnpack(t *testing.T) {
	testCases := []struct {
		testName     string
		input        string
		expectOutput string
		expectErr    error
	}{
		{
			testName:     "default test",
			input:        "a4bc2d5e",
			expectOutput: "aaaabccddddde",
			expectErr:    nil,
		},
		{
			testName:     "string without changing",
			input:        "abcd",
			expectOutput: "abcd",
			expectErr:    nil,
		},
		{
			testName:     "incorrect string",
			input:        "45",
			expectOutput: "",
			expectErr:    ErrIncorrectString,
		},
		{
			testName:     "symbol count > 9",
			input:        "a12",
			expectOutput: "aaaaaaaaaaaa",
			expectErr:    nil,
		},
		{
			testName:     "different digits count",
			input:        "x5y10z",
			expectOutput: "xxxxxyyyyyyyyyyz",
			expectErr:    nil,
		},
		{
			testName:     "correct escape substr",
			input:        `qwe\4\5`,
			expectOutput: "qwe45",
			expectErr:    nil,
		},
		{
			testName:     "correct escape substr",
			input:        `qwe\45`,
			expectOutput: "qwe44444",
			expectErr:    nil,
		},
		{
			testName:     "correct escape substr",
			input:        `qwe\\5`,
			expectOutput: `qwe\\\\\`,
			expectErr:    nil,
		},
		{
			testName:     "incorrect escape substr",
			input:        `qwe\`,
			expectOutput: "",
			expectErr:    ErrIncorrectString,
		},
		{
			testName:     "correct unicode",
			input:        `🐱2🐶3`,
			expectOutput: `🐱🐱🐶🐶🐶`,
			expectErr:    nil,
		},
		{
			testName:     "correct unicode with escape substr",
			input:        `a\\🐱3b`,
			expectOutput: `a\🐱🐱🐱b`,
			expectErr:    nil,
		},
		{
			testName:     "russian unicode",
			input:        "о1ши2бка4",
			expectOutput: "ошиибкаааа",
			expectErr:    nil,
		},
		{
			testName:     "chinese unicode",
			input:        "龙6",
			expectOutput: "龙龙龙龙龙龙",
			expectErr:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := Unpack(tc.input)
			if !errors.Is(err, tc.expectErr) {
				t.Errorf("not equal: expect err: %s, got: %s", tc.expectErr, err)
			}
			if result != tc.expectOutput {
				t.Errorf("not equal string: expect output: %s, got: %s", tc.expectOutput, result)
			}
		})
	}
}

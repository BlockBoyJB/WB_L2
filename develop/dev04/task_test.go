package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	testCases := []struct {
		testName     string
		input        []string
		expectOutput map[string][]string
	}{
		{
			testName: "correct test",
			input:    []string{"тяпка", "пятак", "пятка", "листок", "слиток", "столик"},
			expectOutput: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			testName: "not an anagrams",
			input:    []string{"привет", "мир", "здравствуйте"},
			expectOutput: map[string][]string{
				"привет":       {"привет"},
				"мир":          {"мир"},
				"здравствуйте": {"здравствуйте"},
			},
		},
		{
			testName: "different count anagrams",
			input:    []string{"керамит", "материк", "метрика", "отсечка", "сеточка", "стоечка", "тесачок", "чесотка", "воспрещение", "всепрощение", "просвещение"},
			expectOutput: map[string][]string{
				"керамит":     {"керамит", "материк", "метрика"},
				"отсечка":     {"отсечка", "сеточка", "стоечка", "тесачок", "чесотка"},
				"воспрещение": {"воспрещение", "всепрощение", "просвещение"},
			},
		},
		{
			testName: "repeated words",
			input:    []string{"тяпка", "пятак", "пятка", "листок", "слиток", "столик", "тяпка", "пятак", "пятка", "листок", "слиток", "столик"},
			expectOutput: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			actual := FindAnagrams(tc.input)
			if !reflect.DeepEqual(tc.expectOutput, actual) {
				t.Errorf("not equal: expect %+v, got %+v", tc.expectOutput, actual)
			}
		})
	}
}

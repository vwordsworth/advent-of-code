package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTwoDigit(t *testing.T) {

	testCases := []struct {
		Name     string
		Input    string
		Expected int
	}{
		{
			Name:     "1abc2",
			Input:    "1abc2",
			Expected: 12,
		},
		{
			Name:     "pqr3stu8vwx",
			Input:    "pqr3stu8vwx",
			Expected: 38,
		},
		{
			Name:     "a1b2c3d4e5f",
			Input:    "a1b2c3d4e5f",
			Expected: 15,
		},
		{
			Name:     "treb7uchet",
			Input:    "treb7uchet",
			Expected: 77,
		},
	}
	for _, tt := range testCases {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			result := GetTwoDigits(tt.Input)
			assert.Equal(t, tt.Expected, result)
		})
	}
}

func TestGetTwoDigitWithWords(t *testing.T) {
	testCases := []struct {
		Name     string
		Input    string
		Expected string
	}{
		{
			Name:     "two1nine",
			Input:    "two1nine",
			Expected: "219",
		},
		{
			Name:     "eightwothree",
			Input:    "eightwothree",
			Expected: "8wo3",
		},
		{
			Name:     "abcone2threexyz",
			Input:    "abcone2threexyz",
			Expected: "abc123xyz",
		},
		{
			Name:     "xtwone3four",
			Input:    "xtwone3four",
			Expected: "x2ne34",
		},
		{
			Name:     "4nineeightseven2",
			Input:    "4nineeightseven2",
			Expected: "49872",
		},
		{
			Name:     "zoneight234",
			Input:    "zoneight234",
			Expected: "z1ight234",
		},
		{
			Name:     "7pqrstsixteen",
			Input:    "7pqrstsixteen",
			Expected: "7pqrst6teen",
		},
	}
	for _, tt := range testCases {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			result := ReplaceWordsWithDigits(tt.Input)
			assert.Equal(t, tt.Expected, result)
		})
	}
}

func TestSumList(t *testing.T) {

	testCases := []struct {
		Name     string
		Input    []int
		Expected int
	}{
		{
			Name:     "Non-empty input list",
			Input:    []int{12, 38, 15, 77},
			Expected: 142,
		},
		{
			Name:     "Empty input list",
			Input:    []int{},
			Expected: 0,
		},
	}
	for _, tt := range testCases {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			result := GetListSum(tt.Input)
			assert.Equal(t, tt.Expected, result)
		})
	}
}

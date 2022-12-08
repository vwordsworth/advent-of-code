package bag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBag(t *testing.T) {
	t.Run("Test create new bag", func(t *testing.T) {
		bag := NewBag("vJrwpWtwJgWrhcsFMMfFFhFp")
		assert.Equal(t, "vJrwpWtwJgWr", bag.FirstComp)
		assert.Equal(t, "hcsFMMfFFhFp", bag.SecondComp)
	})
}

func TestGetRepeatedRuneInSingleBag(t *testing.T) {
	type TestCase struct {
		Input    string
		Expected int
	}
	testCases := []TestCase{
		{
			Input:    "vJrwpWtwJgWrhcsFMMfFFhFp",
			Expected: 16,
		},
		{
			Input:    "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			Expected: 38,
		},
		{
			Input:    "PmmdzqPrVvPwwTWBwg",
			Expected: 42,
		},
		{
			Input:    "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			Expected: 22,
		},
		{
			Input:    "ttgJtRGJQctTZtZT",
			Expected: 20,
		},
		{
			Input:    "CrZsJsPPZsGzwwsLwLmpwMDw",
			Expected: 19,
		},
	}

	for _, tc := range testCases {
		t.Run("Test get repeated rune", func(t *testing.T) {
			bag := NewBag(tc.Input)
			assert.Equal(t, tc.Expected, bag.GetSingleBagRepeatedRunePriority())
		})
	}
}

func TestGetRepeatedRuneInBagGroup(t *testing.T) {
	type TestCase struct {
		First    string
		Second   string
		Third    string
		Expected int
	}
	testCases := []TestCase{
		{
			First:    "vJrwpWtwJgWrhcsFMMfFFhFp",
			Second:   "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			Third:    "PmmdzqPrVvPwwTWBwg",
			Expected: 18,
		},
		{
			First:    "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			Second:   "ttgJtRGJQctTZtZT",
			Third:    "CrZsJsPPZsGzwwsLwLmpwMDw",
			Expected: 52,
		},
	}

	for _, tc := range testCases {
		t.Run("Test get repeated rune", func(t *testing.T) {
			group := NewGroup(NewBag(tc.First), NewBag(tc.Second), NewBag(tc.Third))
			assert.Equal(t, tc.Expected, group.GetGroupRepeatedRunePriority())
		})
	}

}

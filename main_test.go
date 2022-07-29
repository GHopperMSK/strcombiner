package strcombiner

import (
	"fmt"
	"testing"
)

func TestCombineSuccess(t *testing.T) {
	dict := []string{"aa", "bb", "abc", "dddd"}

	actualComb := Combine(dict, 0, 0, 9)
	expectedComb := Combination{
		Length: 9,
		Words:  map[int]string{0: "aa", 1: "abc", 2: "dddd"},
	}

	if fmt.Sprint(actualComb.Words) != fmt.Sprint(expectedComb.Words) {
		t.Errorf("Expected \"%s\" is equal to \"%s\"", fmt.Sprint(actualComb.Words), fmt.Sprint(expectedComb.Words))
	}

	if actualComb.Length != expectedComb.Length {
		t.Errorf("Expected length = %d would be equal to %d", actualComb.Length, expectedComb.Length)
	}
}

func TestCombineDecreasedLength(t *testing.T) {
	dict := []string{"aab", "abc", "dddd"}

	actualComb := Combine(dict, 0, 0, 8)
	expectedComb := Combination{
		Length: 7,
		Words:  map[int]string{0: "aab", 1: "dddd"},
	}

	if fmt.Sprint(actualComb.Words) != fmt.Sprint(expectedComb.Words) {
		t.Errorf("Expected \"%s\" is equal to \"%s\"", fmt.Sprint(actualComb.Words), fmt.Sprint(expectedComb.Words))
	}

	if actualComb.Length != expectedComb.Length {
		t.Errorf("Expected length = %d would be equal to %d", actualComb.Length, expectedComb.Length)
	}
}

func TestCombineEmptyDictionary(t *testing.T) {
	dict := []string{}

	actualComb := Combine(dict, 0, 0, 8)
	expectedComb := Combination{
		Length: -1,
		Words:  map[int]string{},
	}

	if fmt.Sprint(actualComb.Words) != fmt.Sprint(expectedComb.Words) {
		t.Errorf("Expected \"%s\" is equal to \"%s\"", fmt.Sprint(actualComb.Words), fmt.Sprint(expectedComb.Words))
	}

	if actualComb.Length != expectedComb.Length {
		t.Errorf("Expected length = %d would be equal to %d", actualComb.Length, expectedComb.Length)
	}
}

func TestCombineTooLongWords(t *testing.T) {
	dict := []string{"1234567890", "abcdefghij"}

	actualComb := Combine(dict, 0, 0, 8)
	expectedComb := Combination{
		Length: -1,
		Words:  map[int]string{},
	}

	if fmt.Sprint(actualComb.Words) != fmt.Sprint(expectedComb.Words) {
		t.Errorf("Expected \"%s\" is equal to \"%s\"", fmt.Sprint(actualComb.Words), fmt.Sprint(expectedComb.Words))
	}

	if actualComb.Length != expectedComb.Length {
		t.Errorf("Expected length = %d would be equal to %d", actualComb.Length, expectedComb.Length)
	}
}

func TestRemoveProcessedSuccess(t *testing.T) {
	dict := []string{"aa", "bb", "abc", "dddd"}
	comb := Combination{
		Length: 9,
		Words:  map[int]string{0: "aa", 1: "abc", 2: "dddd", 3: "c"},
	}

	RemoveProcessed(&dict, comb)

	expectedDict := []string{"bb"}

	if fmt.Sprint(dict) != fmt.Sprint(expectedDict) {
		t.Errorf("Expected \"%s\" is equal to \"%s\"", fmt.Sprint(dict), fmt.Sprint(expectedDict))
	}

	expectedComb := Combination{
		Length: 9,
		Words:  map[int]string{0: "aa", 1: "abc", 2: "dddd", 3: "c"},
	}

	if fmt.Sprint(comb) != fmt.Sprint(expectedComb) {
		t.Errorf("Expected \"%s\" is equal to \"%s\"", fmt.Sprint(comb), fmt.Sprint(expectedComb))
	}
}

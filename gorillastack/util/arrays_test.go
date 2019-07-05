package util

import (
	"reflect"
	"testing"
)

func TestListWithoutDoesNothingWhereNoMatch(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f"}
	result := ListWithout(input, "z")
	if !reflect.DeepEqual(result, input) {
		t.Errorf("ListWithout is wrong. Got: '%s' Expected: '%s'", result, input)
	}
}

func TestListWithoutFiltersWhereMatch(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f"}
	expected := []string{"a", "b", "c", "e", "f"}
	result := ListWithout(input, "d")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ListWithout is wrong. Got: '%s' Expected: '%s'", result, expected)
	}
}

func TestContainsFalse(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f"}
	if result := Contains(input, "z"); result {
		t.Errorf("Contains incorrectly claims \"z\" is in %s", input)
	}
}

func TestContainsTrue(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f"}
	if result := Contains(input, "f"); !result {
		t.Errorf("Contains incorrectly claims \"f\" is not in %s", input)
	}
}

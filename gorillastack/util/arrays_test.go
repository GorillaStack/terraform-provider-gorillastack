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

func TestGetSmallestArrayIndexEmptyList(t *testing.T) {
	input := map[string]interface{}{}
	if result := GetSmallestArrayIndex(input); result != 0 {
		t.Errorf("GetSmallestArrayIndex should return 0 when given an empty array, not %d", result)
	}
}

func createArrayOfActionDefs(index int) []interface{} {
	var actions []interface{}
	return append(actions, map[string]interface{}{"index": index})
}

func helperCreateRawActionsFromIndex(index int) map[string]interface{} {
	rawActions := map[string]interface{}{
		"fake_action_4": createArrayOfActionDefs(index + 3),
		"fake_action_3": createArrayOfActionDefs(index + 2),
		"fake_action_5": createArrayOfActionDefs(index + 4),
		"fake_action_1": createArrayOfActionDefs(index),
		"fake_action_2": createArrayOfActionDefs(index + 1),
	}

	return rawActions
}

func TestGetSmallestArrayIndexWhenIndexedFrom0(t *testing.T) {
	input := helperCreateRawActionsFromIndex(0)
	if result := GetSmallestArrayIndex(input); result != 0 {
		t.Errorf("GetSmallestArrayIndex should return 0 when looking at an array indexed from 0, not %d", result)
	}
}

func TestGetSmallestArrayIndexWhenIndexedFrom1(t *testing.T) {
	input := helperCreateRawActionsFromIndex(1)
	if result := GetSmallestArrayIndex(input); result != 1 {
		t.Errorf("GetSmallestArrayIndex should return 1 when looking at an array indexed from 1, not %d", result)
	}
}

func TestGetSmallestArrayIndexWhenIndexedFrom5(t *testing.T) {
	input := helperCreateRawActionsFromIndex(5)
	if result := GetSmallestArrayIndex(input); result != 5 {
		t.Errorf("GetSmallestArrayIndex should return 5 when looking at an array indexed from 5, not %d", result)
	}
}

package gorillastack

import (
	"testing"

	"github.com/gorillastack/terraform-provider-gorillastack/gorillastack/util"
)

func TestStringArrayOrNullStringWhenNil(t *testing.T) {
	input := StringArrayOrNull{StringArray: nil}
	result := input.String()
	expected := "null"
	if result != expected {
		t.Errorf("StringArrayOrNull is wrong. Got: '%s' Expected: '%s'", result, expected)
	}
}

func TestStringArrayOrNullStringWhenEmptyArray(t *testing.T) {
	input := StringArrayOrNull{StringArray: []*string{}}
	result := input.String()
	expected := "null"
	if result != expected {
		t.Errorf("StringArrayOrNull is wrong. Got: '%s' Expected: '%s'", result, expected)
	}
}

func TestStringArrayOrNullStringWhenEmptyArrayAndShowEmptyTrue(t *testing.T) {
	input := StringArrayOrNull{
		StringArray: []*string{},
		ShowEmpty:   true,
	}
	result := input.String()
	expected := "[]"
	if result != expected {
		t.Errorf("StringArrayOrNull is wrong. Got: '%s' Expected: '%s'", result, expected)
	}
}

func TestStringArrayOrNullStringWithAnArrayOfStrings(t *testing.T) {
	strings := []string{"one", "two", "three"}
	input := StringArrayOrNull{StringArray: util.MapAddresses(strings)}
	result := input.String()
	expected := "[\"one\",\"two\",\"three\"]"
	if result != expected {
		t.Errorf("StringArrayOrNull is wrong. Got: '%s' Expected: '%s'", result, expected)
	}
}

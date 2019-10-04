package util

import (
	"testing"
)

func TestStringValueOnArray(t *testing.T) {
	expected := "[\"one\",\"thing\",0,\"after\",\"another\",false]"
	arr := []interface{}{"one", "thing", 0, "after", "another", false}
	result := StringValue(arr)
	if result != expected {
		t.Errorf("Incorrect array stringification. Got: '%s' Expected: '%s'", result, expected)
	}
}

func TestStringValueOnArrayOfStringPointers(t *testing.T) {
	strings := []string{"one", "two", "three"}
	expected := "[\"one\",\"two\",\"three\"]"
	result := StringValue(MapAddresses(strings))
	if result != expected {
		t.Errorf("Incorrect array stringification. Got: '%s' Expected: '%s'", result, expected)
	}
}

type B struct {
	WhichJeff      string
	IsThisABadTest bool
	SomePrime      int
}

type A struct {
	Name  string
	Bbbbs []*B
}

// test that keys get decapitalised
func TestStringValueOnComplexObject(t *testing.T) {
	input := A{
		Name: "test case",
		Bbbbs: []*B{
			{WhichJeff: "Goldblum", IsThisABadTest: true, SomePrime: 3},
			{WhichJeff: "Garlin", IsThisABadTest: true, SomePrime: 7},
		},
	}
	expected := "{\"name\":\"test case\",\"bbbbs\":[" +
		"{\"whichJeff\":\"Goldblum\",\"isThisABadTest\":true,\"somePrime\":3}," +
		"{\"whichJeff\":\"Garlin\",\"isThisABadTest\":true,\"somePrime\":7}" +
		"]}"
	result := StringValue(input)
	if result != expected {
		t.Errorf("Incorrect array stringification. Got: '%s' Expected: '%s'", result, expected)
	}
}

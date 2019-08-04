package util

import "sort"

func ListWithout(source []string, filterTerm string) []string {
	var filtered []string
	for _, term := range source {
		if term != filterTerm {
			filtered = append(filtered, term)
		}
	}

	return filtered
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func GetSmallestArrayIndex(rawActions map[string]interface{}) int {
	var indecies []int

	for _, rawArrayOfMaps := range rawActions {
		arrayOfMaps := ConvertToArrayOfMaps(rawArrayOfMaps.([]interface{}))

		for _, defn := range arrayOfMaps {
			indecies = append(indecies, defn["index"].(int))
		}
	}

	if len(indecies) == 0 {
		return 0
	}

	sort.Ints(indecies)
	return indecies[0]
}

func ConvertToArrayOfMaps(arr []interface{}) []map[string]interface{} {
	var res []map[string]interface{}
	for _, item := range arr {
		res = append(res, item.(map[string]interface{}))
	}

	return res
}

func MapAddresses(arr []string) []*string {
	LEN := len(arr)
	ptr := make([]*string, LEN)
	for i := 0; i < LEN; i++ {
		ptr[i] = &arr[i]
	}

	return ptr
}

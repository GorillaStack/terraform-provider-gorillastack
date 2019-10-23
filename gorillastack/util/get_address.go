package util

func StringAddress(v string) *string {
	if v == "" {
		return nil
	}

	return &v
}

func BoolAddress(v bool) *bool {
	return &v
}

func IntAddress(i int) *int {
	return &i
}

func MapAddress(m map[string]string) *map[string]string {
	return &m
}

func ArrayOfMapsAddress(m []interface{}) *[]map[string]string {
	var arr []map[string]string

	if len(m) == 0 {
		return nil
	}

	for _, x := range m {
		arr = append(arr, x.(map[string]string))
	}

	return &arr
}

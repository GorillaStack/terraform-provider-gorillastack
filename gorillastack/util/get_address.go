package util

func StringAddress(v string) *string {
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

func ArrayOfMapsAddress(m []map[string]string) *[]map[string]string {
	return &m
}

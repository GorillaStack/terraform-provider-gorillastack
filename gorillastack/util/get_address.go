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

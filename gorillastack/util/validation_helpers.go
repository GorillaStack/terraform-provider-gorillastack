package util

import "regexp"

func GetAwsNamespaceRegex() *regexp.Regexp {
	r, _ := regexp.Compile("/^aws:/")
	return r
}

package util

func GetTagGroupCombiner(v string) *string {
	if v != "" {
		return &v
	}

	return StringAddress("and")
}

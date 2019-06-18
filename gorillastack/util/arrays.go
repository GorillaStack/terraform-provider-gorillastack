package util

func ListWithout(source []string, filterTerm string) []string {
	var filtered []string
	for _, term := range source {
		if term != filterTerm {
			filtered = append(filtered, term)
		}	
	}

	return filtered
}

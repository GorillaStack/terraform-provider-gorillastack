package util

// to be used in cases like accountIds and regions,
// where we want to transform a key of the resourceData to an array of strings. or return nil where not defined
func StringArrayOrNil(v []interface{}) []*string {
	if len(v) > 0 {
		return ArrayOfStringPointers(v)
	}
	
	return nil
}

func ArrayOfStringPointers(arrayOfInterfaces []interface{}) []*string {
	arr := make([]*string, len(arrayOfInterfaces))

	for i, v := range arrayOfInterfaces {
		arr[i] = StringAddress(v.(string))
	}

	return arr
}

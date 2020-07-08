package srutil

// DeduplicateArray : depulicate a slice
func DeduplicateArray(s []string) []string {
	res := NewSet(s...)
	return res.Slice()
}

// ContainsString : return first index of val in array, if val is not in array return -1
func ContainsString(array []string, val string) int {
	for idx, item := range array {
		if item == val {
			return idx
		}
	}
	return -1
}

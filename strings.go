package srutil

// DeduplicateArray : depulicate a slice
func DeduplicateArray(s []string) []string {
	res := NewSet(s...)
	return res.Slice()
}

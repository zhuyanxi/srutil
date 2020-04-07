package srutil

// DeduplicateArray :
func DeduplicateArray(s []string) []string {
	res := NewSet(s...)
	return res.Slice()
}

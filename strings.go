package srutil

// DeduplicateArray :
func DeduplicateArray(s []string) []string {
	// var result []string
	// dedup := make(map[string]struct{})
	// for _, v := range s {
	// 	if _, ok := dedup[v]; !ok {
	// 		dedup[v] = struct{}{}
	// 		result = append(result, v)
	// 	}
	// }

	// return result
	res := NewSet(s...)
	return res.Slice()
}

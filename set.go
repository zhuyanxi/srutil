package srutil

// Set : a set of unique strings
type Set struct {
	m map[string]struct{}
}

// NewSet : create a Set
func NewSet(strings ...string) *Set {
	set := &Set{
		m: map[string]struct{}{},
	}
	for _, s := range strings {
		_ = set.Add(s)
	}
	return set
}

// Add : adds a string to the set. If string is already in the set, it has no effect
// If add successfully, return true
// If already exist, return false
func (s *Set) Add(str string) bool {
	if s.Exist(str) {
		return false
	}
	s.m[str] = struct{}{}
	return false
}

// Exist : checks if string exists in the set
func (s *Set) Exist(str string) bool {
	_, ok := s.m[str]
	return ok
}

// Delete : delete a string from the set
func (s *Set) Delete(str string) {
	delete(s.m, str)
}

// Slice : returns a slice of strings in the set
func (s *Set) Slice() []string {
	n := len(s.m)
	if n == 0 {
		return nil
	}
	arr := make([]string, 0, n)
	for val := range s.m {
		arr = append(arr, val)
	}
	return arr
}

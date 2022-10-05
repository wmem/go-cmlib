package strpkg

import (
	"fmt"
	"sort"
)

// No repeat string array
//
// StringSet is reference type, can directly for function argment and
// should be initialize by NewStringSet()
//
// eg:
//
//	myDB := strpkg.NewStringSet()
type StringSet struct {
	container map[string]byte
}

// New empty StringSet
func NewStringSet() StringSet {
	return StringSet{map[string]byte{}}
}

// Add a string array to Stringset
func (s StringSet) AddArray(array []string) StringSet {
	if array == nil {
		return s
	}

	for _, v := range array {
		s.container[v] = 1
	}
	return s
}

// Convert to string array
func (s StringSet) ToArray() []string {
	array := make([]string, 0, len(s.container))
	for k := range s.container {
		array = append(array, k)
	}
	sort.Strings(array)
	return array
}

// Print element in set
func (s StringSet) Print() {
	for str, _ := range s.container {
		fmt.Println(str)
	}
}

// Add a value to stringset
func (s StringSet) Add(str string) StringSet {
	s.container[str] = 1
	return s
}

// Add other stringset to this stringset
func (s StringSet) AddOther(other StringSet) StringSet {
	for str, _ := range other.container {
		s.container[str] = 1
	}
	return s
}

// Remove a string
func (s StringSet) Remove(str string) StringSet {
	delete(s.container, str)
	return s
}

// Get the length for stringset
func (s StringSet) Length() int {
	return len(s.container)
}

// Is stringset empty
func (s StringSet) IsEmpty() bool {
	return len(s.container) == 0
}

// Is string in stringset
func (s StringSet) IsExist(str string) bool {
	_, exist := s.container[str]
	return exist
}

// Foreach all the element in stringset
func (s StringSet) ForEach(pfnEach func(string)) {
	for str, _ := range s.container {
		pfnEach(str)
	}
}

// Package slices include functions to manipulate slices
package slices

// RemoveDuplicates removes duplicate entries from a slice of string or int
func RemoveDuplicates[T comparable](sliceList []T) []T {
	uniqueList := []T{}
	encountered := make(map[T]bool)

	for _, item := range sliceList {
		if !encountered[item] {
			encountered[item] = true
			uniqueList = append(uniqueList, item)
		}
	}

	return uniqueList
}

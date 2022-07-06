package utils

type any interface{}

const notFoundIndex int = -1

func SliceContains(slice []interface{}, searchElement interface{}) (found bool, indexOfElement int) {
	for index, sliceEl := range slice {
		if sliceEl == searchElement {
			return true, index
		}
	}
	return false, notFoundIndex
}

func SliceContainsString(slice []string, searchElement string) (found bool, indexOfElement int) {
	for index, sliceEl := range slice {
		if sliceEl == searchElement {
			return true, index
		}
	}
	return false, notFoundIndex
}

func FilterSlice(slice []interface{}, operation func(any) bool) []any {
	correctElements := []any{}
	for _, sliceEl := range slice {
		if operation(sliceEl) {
			correctElements = append(correctElements, sliceEl)
		}
	}
	return correctElements
}

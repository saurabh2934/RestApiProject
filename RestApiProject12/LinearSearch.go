package RestApiProject

import "errors"

func LinearSearching(ar []int, find int) error {

	i := 0
	for ; i < len(ar)-1; i++ {
		if find == ar[i] {
			return errors.New("Element Found")
		}
	}
	return errors.New("Element Not Found")
}

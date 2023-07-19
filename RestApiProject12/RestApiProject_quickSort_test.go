package RestApiProject

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	ar := []int{12, 43, 5, 65, 7, 87}
	ar2 := []int{5, 7, 12, 43, 65, 87}
	ar3 := QuickFunction(ar, 0, len(ar)-1)
	if !reflect.DeepEqual(ar3, ar2) {
		t.Fail()
	}
}

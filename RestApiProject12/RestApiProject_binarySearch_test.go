package RestApiProject

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	ar := []int{5, 7, 12, 43, 65, 87}
	find := 2
	er := BinarySearch(ar, find)
	//log.Fatal(er)
	if er.Error() == "Element Not! Found" {
		t.Fail()
	}
}

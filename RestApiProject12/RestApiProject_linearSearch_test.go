package RestApiProject

import "testing"

func TestLinearSearching(t *testing.T) {
	ar := []int{12, 43, 5, 65, 7, 87}
	er := LinearSearching(ar, 4)

	if er.Error() == "Element Not Found" {
		t.Fail()
	}
}

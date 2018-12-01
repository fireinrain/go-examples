package qsort

import "testing"

func TestQuickSort(t *testing.T) {
	values := []int{5, 3, 6, 1, 0, 4}

	QuickSort(values)

	if values[0] != 0 || values[1] != 1 || values[2] != 3 || values[3] != 4 || values[4] != 5 || values[5] != 6 {
		t.Error("quicksort error got: ", values, "excepted 0 1 3 4 5 6")
	}
}

package bubblesort

import "testing"

func TestBubblesort(t *testing.T) {
	values := []int{5, 4, 3, 21, 6}

	Bubblesort(values)
	if values[0] != 3 || values[1] != 4 || values[2] != 5 || values[3] != 6 || values[4] != 21 {
		t.Error("BubbleSort() failed .Got: ", values, "Excepted: 3 4 5 6 21")
	}
}

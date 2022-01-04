package internal

import (
	"testing"
)

func Test_CountOfSearch(t *testing.T) {
	list := make([]int, 100)
	for i := 0; i < len(list); i++ {
		list[i] = i + 1
	}
	if list[0] != 1 {
		t.Errorf("the first element need be 1")
	}
	if list[len(list)-1] != 100 {
		t.Errorf("the last element need be 100")
	}

	t.Run("item = 1", func(t *testing.T) { checkResult(t, list, 1, 0, 6) })
	t.Run("item = 100", func(t *testing.T) { checkResult(t, list, 100, 99, 7) })
	t.Run("item = 50", func(t *testing.T) { checkResult(t, list, 50, 49, 1) })
	t.Run("item = 63", func(t *testing.T) { checkResult(t, list, 63, 62, 6) })
}

func checkResult(t *testing.T, list []int, item, needIndex, needCount int) {
	index, count, err := BinarySearch(list, item)
	if err != nil {
		t.Fatal(err)
	}
	if index != needIndex {
		t.Errorf("item need be at %d element, but have: %d", needIndex, index)
	}
	if count != needCount {
		t.Errorf("count need be %d, but have: %d", needCount, count)
	}
}

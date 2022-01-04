package internal

import (
	"errors"
)

func BinarySearch(list []int, item int) (int, int, error) {
	high := len(list) - 1
	count := 0

	for low := 0; low <= high; {
		count += 1

		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return mid, count, nil
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, count, errors.New("not found")
}

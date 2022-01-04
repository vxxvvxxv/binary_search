package main

import (
	"fmt"

	"github.com/vxxvvxxv/binary_search/internal"
)

func main() {
	list := []int{1, 3, 5, 7, 9}

	item, _, err := internal.BinarySearch(list, 3)
	fmt.Println(list, item, err)

	item, _, err = internal.BinarySearch(list, -1)
	fmt.Println(list, item, err)
}

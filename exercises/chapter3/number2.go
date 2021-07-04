package chapter3

import "fmt"

func bubbleSort(slice *[]int) {
	for i := len(*slice) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if (*slice)[j] > (*slice)[j+1] {
				(*slice)[j+1], (*slice)[j] = (*slice)[j], (*slice)[j+1]
			}
		}
	}
}

func SortSlice() {
	slice := []int{4, 3, 2, 6, 4, 33, 10, 32, 15, 38, 24, 1, 3}
	bubbleSort(&slice)
	fmt.Println("Sorted array")
	for i := range slice {
		fmt.Printf("%d ", slice[i])
	}
}

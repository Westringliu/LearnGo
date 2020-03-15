package main

import "fmt"

func oddEvenSort(arr []int) []int {
	isSorted := false

	for isSorted == false {
		isSorted = true
		for i := 1; i < len(arr)-1; i = i + 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		fmt.Println("o", arr)
		for i := 0; i < len(arr)-1; i = i + 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		fmt.Println("e", arr)
	}

	return arr
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	fmt.Println(oddEvenSort(arr))
}

package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		newPos := find(arr, i, arr[i])
		shift(arr, newPos, i)
	}
}

func find(arr []int, end int, element int) (pos int) {
	for i := 0; arr[i] < element && i <= end; i++ {
		pos = i + 1
	}
	return
}

func shift(arr []int, start int, end int) {
	for i := end; i > start; i-- {
		arr[i], arr[i-1] = arr[i-1], arr[i]
	}
}

func main() {
	arr := []int{9, 2, 3, 4, 5, 6, 7, 8, 1}
	insertionSort(arr)
	fmt.Println(arr)
}

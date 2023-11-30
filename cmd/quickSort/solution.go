package main

import "fmt"

func getPivot(arr []int) int {
	return arr[len(arr)/2]
}

func quickSort(arr []int, p func([]int) int) {
	quickSortInner(&arr, p, 0, len(arr)-1)
}

func quickSortInner(arr *[]int, p func([]int) int, start int, end int) {
	if end-start < 4 {
		insertSort(arr, start, end)
		return
	}

	var first, last = start, end
	pivot := (*arr)[(end+start)/2]
	for first <= last {
		for (*arr)[first] < pivot && first < end {
			first++
		}
		for (*arr)[last] > pivot && last > start {
			last--
		}

		if first <= last {
			(*arr)[last], (*arr)[first] = (*arr)[first], (*arr)[last]
			first++
			last--
		}
	}
	if start < first {
		quickSortInner(arr, p, start, first-1)
	}
	if end > last {
		quickSortInner(arr, p, last+1, end)
	}
}

func insertSort(arr *[]int, start int, end int) {
	for i := start; i <= end; i++ {
		newPos := find(arr, i, (*arr)[i])
		shift(arr, newPos, i)
	}
}

func find(arr *[]int, end int, element int) (pos int) {
	for i := 0; (*arr)[i] < element && i <= end; i++ {
		pos = i + 1
	}
	return
}

func shift(arr *[]int, start int, end int) {
	for i := end; i > start; i-- {
		(*arr)[i], (*arr)[i-1] = (*arr)[i-1], (*arr)[i]
	}
}

func main() {
	data := []int{1}
	quickSort(data, getPivot)
	fmt.Println(data)
}

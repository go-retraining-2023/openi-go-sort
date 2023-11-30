package main

import "fmt"

func countingSort(arr []int) {
	min, max := count(arr)
	counts := make([]int, max-min+1)
	for _, v := range arr {
		counts[v-min]++
	}

	ind := 0
	for i := 0; i < len(counts); i++ {
		for j := 0; j < counts[i]; j++ {
			arr[ind] = i + min
			ind++
		}
	}
	fmt.Println(arr)
}

func count(arr []int) (int, int) {
	min, max := 0, 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < arr[min] {
			min = i
		}
		if arr[i] > arr[max] {
			max = i
		}
	}

	if len(arr) == 0 {
		return 0, 0
	}
	return arr[min], arr[max]
}

func main() {
	arr := []int{1, 2, 2, 2, 1, 0, 0, 3, 15, 2}
	countingSort(arr)
}

package main

import "fmt"

func mergeSort(arr []int) {
	mergeSortInner(arr, 0, len(arr)-1)
}

func mergeSortInner(arr []int, start int, end int) {
	if end-start < 1 {
		return
	} else if end-start == 1 {
		if arr[start] > arr[end] {
			arr[start], arr[end] = arr[end], arr[start]
		}
	} else {
		mergeSortInner(arr, start, (end+start)/2)
		mergeSortInner(arr, (start+end)/2+1, end)
		merge(arr, start, (end+start)/2, (end+start)/2+1, end)
	}
}

func merge(arr []int, s1 int, e1 int, s2 int, e2 int) {
	tmp := make([]int, e1-s1+e2-s2+2)
	i, j, k := 0, s1, s2
	for i < len(tmp) && j <= e1 && k <= e2 {
		if arr[j] < arr[k] {
			tmp[i] = arr[j]
			j++
		} else {
			tmp[i] = arr[k]
			k++
		}
		i++
	}

	if j > e1 {
		for k <= e2 {
			tmp[i] = arr[k]
			i++
			k++
		}
	} else {
		for j <= e1 {
			tmp[i] = arr[j]
			i++
			j++
		}
	}

	j = 0
	for i = s1; i <= e1; i++ {
		arr[i] = tmp[j]
		j++
	}
	for i = s2; i <= e2; i++ {
		arr[i] = tmp[j]
		j++
	}
}

func main() {
	data := []int{4, 5, 4, 9, 8, 7, 15, 12}
	mergeSort(data)
	fmt.Println(data)
}

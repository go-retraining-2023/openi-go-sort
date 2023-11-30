package main

import "fmt"

var min, max int

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

func msBits(arr []int, item int) int {
	return (item - min) / (max - min + 1) * len(arr)
}

func bucketSort(arr []int, f func([]int, int) int) {
	if len(arr) < 1 {
		return
	}
	var minI, maxI = 0, 0
	for i := 1; i < len(arr); i++ {
		if arr[minI] > arr[i] {
			minI = i
		}
		if arr[maxI] < arr[i] {
			maxI = i
		}
	}
	min = arr[minI]
	max = arr[maxI]
	var buckets map[int][]int = make(map[int][]int)
	for _, item := range arr {
		buckets[f(arr, item)] = append(buckets[f(arr, item)], item)
	}

	j := 0
	for _, v := range buckets {
		insertionSort(v)
		for i := 0; i < len(v); i++ {
			arr[j] = v[i]
			j++
		}
	}
}

func main() {
	data := []int{}
	bucketSort(data, msBits)
	fmt.Println(data)
}

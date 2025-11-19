package main

import "fmt"

func main() {
	list := []int{5, 3, 6, 2, 10}
	// item := 3
	// fmt.Println(binarySearch(list, item))
	// fmt.Println(selectionSort(list))
	// countdown(3)
	//
	// fmt.Println(fact(3))
	// fmt.Println(sum(list))
	// fmt.Println(countElements(list))
	fmt.Println(quickSort(list))
}

func binarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high)
		if list[mid] == item {
			return mid
		} else if list[mid] > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallest_index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func selectionSort(arr []int) []int {
	newArr := []int{}
	// Лучше использовать цикл while-стиля (пока массив не пуст)
	for len(arr) > 0 {
		smallest := findSmallest(arr)
		newArr = append(newArr, arr[smallest])
		// Исправлено: добавлено ":" и многоточие
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return newArr
}

func countdown(i int) {
	if i < 1 {
		return
	} else {
		fmt.Println(i)
		countdown(i - 1)
	}
}

func fact(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * fact(x-1)
	}
}

func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else {
		return nums[0] + sum(nums[1:])
	}
	// return 1919
}

func countElements(arr []int) int {
	if len(arr) == 1 {
		return 1
	} else if len(arr) == 0 {
		return 0
	} else {
		return 1 + countElements(arr[1:])
	}
}

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	} else {
		pivot := array[0]
		less := []int{}
		for _, value := range array[1:] {
			if value <= pivot {
				less = append(less, value)
			}
		}
		greater := []int{}
		for _, value := range array[1:] {
			if value > pivot {
				greater = append(greater, value)
			}
		}
		result := append(quickSort(less), pivot)
		result = append(result, quickSort(greater)...)
		return result
	}
}

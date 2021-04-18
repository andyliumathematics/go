package main

import "fmt"

func main() {
	// var arr = [...]int64{3, 4, 5, 6, 7, 8, 34, 5, 446, 4}
	// fmt.Println(arr)
	// for firstposition := 0; firstposition < len(arr)-1; firstposition++ {
	// 	minposition := firstposition
	// 	for i := firstposition + 1; i < len(arr); i++ {
	// 		if arr[minposition] > arr[i] {
	// 			minposition = i
	// 		}
	// 	}
	// 	if minposition != firstposition {
	// 		tmp := arr[minposition]
	// 		for j := minposition; j > firstposition; j-- {
	// 			arr[j] = arr[j-1]
	// 		}
	// 		arr[firstposition] = tmp
	// 	}
	// }
	// fmt.Println(arr)
	selectsort2()
}
func selectsort2() {
	var arr = [...]int64{23, 84, 15, 36, 7, 8, 34, 5, 446, 4}
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if i != min {
			tmp := arr[i]
			arr[i] = arr[min]
			arr[min] = tmp
		}

	}
	fmt.Println(arr)

}

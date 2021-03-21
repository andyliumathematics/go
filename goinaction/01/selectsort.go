// +build ignore

package main

import "fmt"

func main() {
	var arr = [...]int64{3, 4, 5, 6, 7, 8, 34, 5, 446, 4}
	fmt.Println(arr)
	for firstposition := 0; firstposition < len(arr)-1; firstposition++ {
		minposition := firstposition
		for i := firstposition + 1; i < len(arr); i++ {
			if arr[minposition] > arr[i] {
				minposition = i
			}
		}
		if minposition != firstposition {
			tmp := arr[minposition]
			for j := minposition; j > firstposition; j-- {
				arr[j] = arr[j-1]
			}
			arr[firstposition]=tmp
		}
	}
	fmt.Println(arr)
}

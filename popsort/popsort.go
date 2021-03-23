package main

import "fmt"

func main() {
	var arr = [...]int64{8, 4, 5, 6, 7, 8, 34, 5, 446, 4,1}

	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	fmt.Print(arr)
}

package main

import "fmt"

func main() {
	var arr = [...]int64{8, 4, 5, 6, 7, 8, 34, 5, 446, 4}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] >= arr[i] {
				tmp := arr[i]
				for k := i - 1; k >= j; k-- {
					arr[k+1] = arr[k]
				}
				arr[j] = tmp
				break
			}
		}
	}
	fmt.Println(arr)
	sort2()

}
func sort2() {
	var arr = [...]int64{8, 4, 5, 6, 7, 8, 34, 5, 446, 4}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j+1] < arr[j] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			} else {
				break
			}
		}
	}
	fmt.Print(arr)

}

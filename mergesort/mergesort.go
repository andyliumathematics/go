package main

import "fmt"

func main() {
	var arr = []int{8, 4, 5, 6, 7, 8, 34, 5, 446, 4, 1}
	fmt.Println(merge(arr))
	// slise := make(arr, len(arr)/2)

}
func merge(data []int) []int {
	var arr []int
	if len(data) > 1 {
		left := merge(data[0 : len(data)/2])
		right := merge(data[len(data)/2:])
		m := 0
		l := 0
		j := 0
		for ; j < len(arr) && m < len(left) && l < len(right); j++ {
			if left[m] < right[l] {
				arr[j] = left[m]
				m++
			} else {
				arr[j] = right[l]
				l++
			}
		}
		if m < len(left) {
			for k := m; k < len(left); k++ {
				arr[j] = left[k]
				j++
			}
		} else if l < len(right) {

			for k := l; k < len(right); k++ {
				arr[j] = right[k]
			}
		}
	}
	return arr
}

package main

import "fmt"

func main() {
	var arr = [...]int64{3, 4, 5,33,63, 6, 7, 8, 34, 5, 446, 4}
	for i := 1; i < len(arr); i++ {
		//位置
		tmp := arr[i]
		for j := i-1; j >= 0; j-- {
			if arr[j] >= tmp {
				// arr[j+1] = arr[j]
				// arr[j] = tmp
				swap(&arr,j,j+1)
			}
		}
	}
	fmt.Println(arr)
}
func swap(arr *[12]int64,i int,j int){
	tmp := arr[i]
	arr[i]=arr[j]
	arr[j] = tmp
}
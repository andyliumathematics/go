package main

import "fmt"

// import "fmt"

func main() {
	// var arr = [...]int64{23, 84, 15, 36, 7, 8, 34, 5, 446, 4}

	shellSort3()
	// for i := len(arr) / 2; i > 0; i /= 2 {
	// 	for k := 0; k < i; k++ {
	// 		for j := k; j < len(arr); j += i {
	// 			if j+i < len(arr) && arr[j] > arr[j+i] {
	// 				tmp := arr[j]
	// 				arr[j] = arr[j+i]
	// 				arr[j+i] = tmp
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Print(arr)
}
func shellSort3() {
	var arr = [...]int64{23, 84, 15, 36, 7,7, 8, 34, 5, 446, 4,99}
	for i := len(arr) / 2; i > 0; i /= 2 { //gap不断缩小
		for j := 0; j < i; j++ { //每个GAP 每一组数据都循环一遍
			for k := j + i; k < len(arr); k += i {
				for l := k; l >= i; l -= i {
					if arr[l] < arr[l-i] {
						tmp := arr[l]
						arr[l] = arr[l-i]
						arr[l-i] = tmp
					} else {
						break
					}
				}

			}

		}
	}
	fmt.Println(arr)

}

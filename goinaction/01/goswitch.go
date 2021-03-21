// +build ignore

package main

import (
	"fmt"
)
type xx struct{
	a int
	xx 
}
func main33333(){
	x := "d"
	switch{
	case x=="d":
		fmt.Print("ddd")
		fallthrough
	case x== "":
		fmt.Print("   ....  .")
		fallthrough
	case x=="p":
		fmt.Print("ppppppp")
		fallthrough
	case x=="t":
		fmt.Print("tttttttttt")
		fallthrough
	case x=="b":
		fmt.Print("bbbbbbbbbbbbbbb")
		fallthrough
	default:
		fmt.Print("default")

		aa := xx{
			a:3,
			xx{
				a:87,
				xx:nil
			}
		}
		fmt.Println(aa )
	}
}
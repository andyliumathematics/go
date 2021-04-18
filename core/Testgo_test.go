package main

import (
	"encoding/json"
	"fmt"
	"testing"
)
type X struct{
	A int
	B int
	C *X
}
func TestGo00(t *testing.T) {
	ax := X{}
	ax.A=23
	ax.B = 2333
	ax.C = &X{A:32,B:3,C:nil}
	a,err :=json.Marshal(ax)
	fmt.Println(err)
	fmt.Println(string(a))
	// fmt.Print("%V",ax.c)

}

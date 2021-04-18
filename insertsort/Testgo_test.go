package main

import (
	"fmt"
	"testing"
)
type X struct{
	a int
	b int
	c *X
}
func TestGo00(t *testing.T) {
	ax := X{}
	ax.a=23
	ax.b = 2333
	ax.c = &X{a:32,b:3,c:nil}
	fmt.Print("%V",ax.c)

}

package main

import (
	"fmt"

	"github.com/sagar23sj/greeting"
	"github.com/sagar23sj/hello/num"
)

func main() {
	str := greeting.Hello("")
	fmt.Println("String Val : ", str)

	ret := num.Add(1, 3)
	fmt.Println("Return val : ", ret)
}

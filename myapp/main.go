package main

import (
	"fmt"

	"github.com/rajeshpillai/rockety-go"
)

func main() {
	result := rockety.TestFunc(1, 1)
	fmt.Println(result)

	result = rockety.TestFunc2(4, 1)
	fmt.Println(result)

	result = rockety.TestFunc3(2, 2)
	fmt.Println(result)
}

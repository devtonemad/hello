package main

import (
	"fmt"

	"github.com/devtonemad/stringutil"
)

func main() {

	fmt.Printf("Hello World!\n")

	fmt.Printf(stringutil.Reverse("This is a reversed Message! Can you read?"))

	fmt.Printf("\n")

	loop()

}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("counter:%v\n", i)
	}
}

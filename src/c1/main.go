package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("golang-learning")
	fmt.Println(os.Args)
	fmt.Println(len(os.Args))

	var args = os.Args

	for a := 0; a < len(args); a++ {
		fmt.Println(args[a])
	}

	os.Exit(1)
}

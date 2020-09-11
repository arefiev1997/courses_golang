package main

import (
	"fmt"

	"github.com/arefiev1997/courses_golang/internal/lesson2"
)

func main() {
	fmt.Println("Starting lesson 2")
	lesson2.Run()

	lesson2.RunFunctions()
	lesson2.RunInterface(5)

	fmt.Println("End of lesson 2")
}

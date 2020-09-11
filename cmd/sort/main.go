package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args[0])
	sortType := os.Args[1]

	// switch sortType {
	// case "puz":
	// 	funcPuz()
	// default:

	// }
	fmt.Println(sortType)

	inputArray := make([]int, len(os.Args[2:]))
	var err error
	for i, v := range os.Args[2:] {
		inputArray[i], err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(inputArray)
}

package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Subslice changing subslice
func Subslice() {
	s := []int{1, 2, 3}
	ss := s[1:]
	// ss = append(ss, 4)

	for _, v := range ss {
		v += 10
	}

	for i := range ss {
		ss[i] += 10
	}

	fmt.Println(s)
}

// DeclareVariables declaring new variables
func DeclareVariables() {

	var x int

	// x, _ := getVariables()
	// x, _ = getVariables()
	x, y := getVariables()
	// x, y = getVariables()

	fmt.Printf("%d %d", x, y)
}

func getVariables() (int, int) {
	return 1, 2
}

// MapOK check key in map
func MapKey() {
	m := make(map[string]int) //A
	m["b"] = 1
	if v, ok := m["b"]; ok { //B
		fmt.Println(v)
	}
}

// LoopPointer loop
func LoopPointer() {
	m := make(map[int]*int)

	N := 3

	for i := 0; i < N; i++ {
		k := i
		m[i] = &k //A
	}

	for _, v := range m {
		fmt.Print(*v)
	}
	fmt.Println()
}

// BreakOuter break outer
func BreakOuter() {
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(i, ",", j, " ")
			break outer
		}
		fmt.Println()
	}
}

func OpenFile() {
	f, err := os.Open("file")
	defer f.Close()
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(f)
	println(string(b))
}

func F1() {
	defer println("f1-begin")
	f2()
	defer println("f1-end")
}

func f2() {
	defer println("f2-begin")
	f3()
	defer println("f2-end")
}

func f3() {
	defer println("f3-begin")
	panic(0)
	defer println("f3-end")
}

func Recov() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recover:%#v", r)
		}
	}()
	panic(1)
	panic(2)
}

type Result struct {
	Status int
}

func JsonUnmarshall() {
	var data = []byte(`{"status": 200}`)
	result := &Result{}

	if err := json.Unmarshal(data, result); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("result=%+v", result)
}

// func RunGoroutines() {
// 	wg := &sync.WaitGroup{}
// 	go func() {
// 		fmt.Println("text")
// 	}
// }

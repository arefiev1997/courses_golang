package tasks

import "fmt"

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

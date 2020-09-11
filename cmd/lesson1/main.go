package main

import (
	"errors"
	"fmt"
)

var globalValue *int

const constantValue = 10

func createPointer() {
	k := 5
	globalValue = &k
	fmt.Println(k)
}

func main() {

	// Инициализация переменных
	// целочисленные типы
	var intValue8 int8        // используют
	var intValue16 = 12       // не используют
	intValue32 := 16          // используют
	var intValue64 int64 = 23 // не используют
	// intValue64V2 := int64(23)

	// var intValue int

	// var uintValue8 uint8

	intValue32 = 17
	fmt.Printf("%d %d %d %d\n", intValue8, intValue16, intValue32, intValue64)

	// loopUint()

	// var floatValue float32
	// var floatValu64 float64

	// fmt.Println(float64(10) / float64(7))

	// Строковый тип
	s := "Hello"

	fmt.Printf("%c\n ", s[0])

	// Булевый тип
	boolValue := false

	if !boolValue {
		fmt.Println("false ")
	}

	// Указатели
	i := 10
	p := &i
	*p = 11
	fmt.Println(i)

	// Обработка ошибки при делении на 0
	divideResult, err := divide(10, 1)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(divideResult)
	}

	if divideResult, err = divide(10, 1); err != nil {
		panic(err)
	} else {
		fmt.Println(divideResult)
	}

	intValue1 := new(int)
	handlePointer(intValue1)

	// Массивы
	var array [3]int
	fmt.Println(array)

	array[1] = 5
	for _, v := range array {
		fmt.Printf("%d\n", v)
	}

	changeArray(array)
	fmt.Println(array)

	// Срезы
	sliceValue := array[1:]
	fmt.Println(sliceValue)

	fmt.Printf("len - %d, cap - %d\n", len(sliceValue), cap(sliceValue))

	sliceValue = append(sliceValue, 1, 2, 3, 4)
	fmt.Printf("len - %d, cap - %d\n", len(sliceValue), cap(sliceValue))

	sliceValue[0] = 150
	fmt.Println(sliceValue)
	fmt.Println(array)

	sliceValue3 := make([]int, len(array))
	for i, v := range array {
		sliceValue3[i] = v + 10
	}
	fmt.Println(sliceValue3)

	// map
	ages := map[string]int{
		"vasya": 1,
	}

	handleMap(ages)
	fmt.Println(ages)

	fmt.Println(array)

	// Изменение массива по адресу
	changeArrayByPointer(&array)
	fmt.Println(array)

	// Вывод всех ключей-значений мапа
	for i, v := range ages {
		fmt.Printf("key - %s, value - %d\n", i, v)
	}

	ages["len"]++
	fmt.Println(ages["len"])

	if v, ok := ages["ken"]; ok {
		fmt.Println(v)
	}

	var defaultMap map[string]int
	defaultMap = make(map[string]int)
	defaultMap = map[string]int{}

	fmt.Println(defaultMap["l"])
	defaultMap["le"] = 10
	fmt.Println(defaultMap)
}

func handleMap(m map[string]int) {
	m["lena"] = 14
}

func changeArrayByPointer(arrayToChange *[3]int) {
	arrayToChange[0] = 11
}

func changeSlice(sliceValue []int) []int {
	sliceValue2 := append(sliceValue, 1, 2, 3)
	return sliceValue2
}

// бесконечный цикл
func loopUint() {
	for i := uint(4); i >= 0; i-- {
		fmt.Println(i)
	}
}

func handleIntegerPointer(intValue *int) {

	if intValue == nil {
		return
	}

}

func divide(value1, value2 int) (int, error) {
	if value2 == 0 {
		return 0, errors.New("Dividing by zero")
	}

	return value1 / value2, nil
}

func handlePointer(pointer *int) {
	fmt.Println(pointer)
}

func changeArray(ar [3]int) {
	ar[0] = 10
}

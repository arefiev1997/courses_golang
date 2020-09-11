package lesson2

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// RunFunctions run functions
func RunFunctions() error {
	// Обработка аварийной ситуации. анонимная функция вызовется перед выходом из функции RunFunctions
	defer func() error {
		if r := recover(); r != nil {
			fmt.Println(r)
			return fmt.Errorf("%v", r)
		}
		return nil
	}()
	x, y := 10, 0

	// тут происходит попытка подключиться к бд. код можно закомментить
	db, err := sqlx.Connect("driverName string", "re")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	compareResult := compareFuncsResult(x, y, multiply, divide)
	

	fmt.Println(compareResult)

	variadic("1", "2", "3")
	return nil
}

// Передача функций в качестве аргументов. Тут мы сравниваем результаты двух функций
func compareFuncsResult(x, y int, f1 func(int, int) int, f2 func(int, int) int) bool {
	return f1(x, y) > f2(x, y)
}

func divide(x, y int) int {
	return x / y
}

func multiply(x, y int) int {
	return x * y
}

func variadic(args ...string) {
	fmt.Println(args)
}

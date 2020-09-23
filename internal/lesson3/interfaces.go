package lesson3

import "fmt"

type Printer interface {
	PrintHello()
}

type RussianPrinter struct {
	msg string
}

func (r *RussianPrinter) PrintHello() {
	fmt.Println("Привет")
}

type EnglishPrinter struct {
}

func (e *EnglishPrinter) PrintHello() {
	fmt.Println("Hello")
}

func GetRussianPrinter() *RussianPrinter {
	return nil
}

func HandlePrinter(printer Printer) {
	if printer == nil {
		return 
	}
}
package main

import "github.com/arefiev1997/courses_golang/internal/lesson3"

func main() {
	// var p *lesson3.RussianPrinter
	// p = lesson3.GetRussianPrinter()
	// fmt.Println(p)
	// if p == nil {
	// 	fmt.Println("NIL")
	// }

	// wg := &sync.WaitGroup{}
	// wg.Add(1)
	// go lesson3.SleepPrint(wg)
	// fmt.Println("hello")
	// fmt.Println(runtime.Stack())
	// wg.Wait()
	// wg.Add(1)
	// wg.Wait()
	// fmt.Println("end of sleep")

	lesson3.RunReadAndPrintFile()
}

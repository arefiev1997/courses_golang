package lesson4

import (
	"fmt"
	"runtime"
	s "sync"
	"time"
)

var mp = map[int]string{}
var t int
var mu s.Mutex
var rw s.RWMutex
var c chan struct{}

const (
	Monday int = iota
	Tuesday
)

func increase() {
	t++
}

// Run run
func Run() {

	mu = s.Mutex{}
	rw = s.RWMutex{}

	fmt.Println(Monday)
	fmt.Println(Tuesday)

	c = make(chan struct{})
	workersCount := 10
	wg := &s.WaitGroup{}
	wg1 := &s.WaitGroup{}

	// wg1.Add(1)
	// go increaseT(wg1)

	start := time.Now()
	wg.Add(workersCount)
	for i := 0; i < workersCount; i++ {
		go setMapValue(wg)
		// go getValue(wg)
	}
	fmt.Printf("GoroutinesCount - %d\n", runtime.NumGoroutine())
	wg.Wait()
	close(c)

	wg1.Wait()
	fmt.Printf("Execution time - %s\n", time.Since(start).String())
}

// func setMapValue(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 10; i++ {
// 		t++
// 		fmt.Println(t)
// 	}
// }

// func setMapValue(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 10; i++ {
// 		// value := rand.Intn(100)
// 		c <- struct{}{}

// 		// mp[value] = strconv.Itoa(value)
// 	}
// }

// func increaseT(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for range c {
// 		t++
// 		fmt.Println(t)
// 	}
// }

// func setMapValue(wg *s.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 100; i++ {
// 		mu.Lock()
// 		t++
// 		fmt.Println(t)
// 		mu.Unlock()
// 	}
// }

// func getValue(wg *s.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 100; i++ {
// 		mu.Lock()
// 		fmt.Println(t)
// 		mu.Unlock()
// 	}
// }

// func setMapValue(wg *s.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(t)
// 	}
// }

func setMapValue(wg *s.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1; i++ {
		rw.Lock()
		t++
		fmt.Println(t)
		rw.Unlock()
	}
}

func getValue(wg *s.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1; i++ {
		rw.RLock()
		fmt.Println(t)
		rw.RUnlock()
	}
}

// func setMapValue() {
// 	atomic.AddInt64(&t, 1)
// }

package lesson3

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func SleepPrint(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(time.Second * 1))
	fmt.Println("sleep")
}

func ReadFile(ctx context.Context, wg *sync.WaitGroup, fileString chan string) error {
	defer func() {
		close(fileString)
		wg.Done()
	}()
	file, err := os.Open("/Users/artemarefev/golang_cources/internal/lesson2/structs.go")
	if err != nil {
		return err
	}
	defer file.Close()

	f := bufio.NewReader(file)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		line, err := f.ReadString('\n')
		if err == io.EOF {
			fmt.Println("EOF")
			break
		} else if err != nil {
			fmt.Println(err)
			return nil
		}
		fileString <- line
		// fmt.Println(line)
	}
	fmt.Println("END")
	return nil
}

func PrintStringFromFile(ctx context.Context, wg *sync.WaitGroup, fileString <-chan string, workerNum int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("DONE")
			return
		case s := <-fileString:
			if s == "" {
				return
			}
			fmt.Printf("workerNum - %d, string - %s\n", workerNum, s)
			time.Sleep(time.Duration(time.Second * 1))
		}
	}
}

func PrintStringFromFileV2(ctx context.Context, wg *sync.WaitGroup, fileString <-chan string, workerNum int) {
	defer wg.Done()
	for s := range fileString {
		fmt.Printf("workerNum - %d, string - %s\n", workerNum, s)
		time.Sleep(time.Duration(time.Second * 1))
	}
}

func RunReadAndPrintFile() {
	fileStrings := make(chan string)

	wg := &sync.WaitGroup{}
	ctx := context.Background()
	erg, ctx := errgroup.WithContext(ctx)
	ctx, cancel := context.WithCancel(ctx)
	wg.Add(2)
	var workersCount = 1

	erg.Go(func() error {
		return ReadFile(ctx, wg, fileStrings)
	})
	// go ReadFile(wg, fileStrings)
	for i := 0; i < workersCount; i++ {
		go PrintStringFromFileV2(ctx, wg, fileStrings, i)
	}
	time.Sleep(time.Duration(2 * time.Second))
	cancel()
	// wg.Wait()
	err := erg.Wait()
	if err != nil {
		fmt.Println(err)
	}
}

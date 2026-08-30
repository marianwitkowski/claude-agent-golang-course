package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTask() <-chan int32 {
	ch := make(chan int32)

	go func() {
		defer close(ch)
		time.Sleep(time.Millisecond * 2500)
		ch <- rand.Int31n(123)
	}()
	return ch
}

func printLetters(ch chan<- string) {
	for i := 'a'; i <= 'z'; i++ {
		ch <- fmt.Sprintf("%c", i) // fmt.Printf("%c ", i)
		time.Sleep(time.Millisecond * 30)
	}
	close(ch)
}

func printNumbers(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i //fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 60)
	}
	close(ch)
}

func main() {

	//v1 := <-longTask()
	//fmt.Println(v1)

	ch1, ch2, ch3 := longTask(), longTask(), longTask()
	a, b, c := <-ch1, <-ch2, <-ch3
	fmt.Println(a, b, c)

	lettersCh := make(chan string)
	numbersCh := make(chan int)

	go printLetters(lettersCh)
	go printNumbers(numbersCh)

	//time.Sleep(time.Second * 1)
	for {
		// tu odbieramy dane
		select {

		case letter, ok := <-lettersCh:
			if ok {
				fmt.Printf("%s ", letter)
			} else {
				lettersCh = nil
			}

		case number, ok := <-numbersCh:
			if ok {
				fmt.Printf("%d ", number)
			} else {
				numbersCh = nil
			}
		}

		// sprawdzanie czy kanały zamknięte
		if lettersCh == nil && numbersCh == nil {
			break
		}
	}
	fmt.Println("\nKoniec programu")
}

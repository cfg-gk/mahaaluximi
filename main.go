package main

import (
	"fmt"
	"time"
)

func genNumbers(ch chan int, n int) {
	ch <- n
}

func primeNumberValidate(ch chan int) {
	count := 0
	number := <-ch

	for i := 2; i <= number/2; i++ {

		if number%i == 0 {
			count++
		}

	}
	if count == 0 && number != 1 {
		fmt.Println(number, "-> Prime Number")
	} else {
		fmt.Println(number, "-> Not Prime Number")
	}

}

func main() {
	num := make(chan int)
	loop := 1
	for loop <= 100 {
		go genNumbers(num, loop)
		go primeNumberValidate(num)
		loop++
		time.Sleep(2 * time.Millisecond)
	}
}

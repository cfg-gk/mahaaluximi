package main

import (
	"fmt"
	"strconv"
	"time"
)

func genAlphabets(c chan<- string, n string) {
	//num := 65
	//var alphabet string
	// for i := 65; i <= 90; i++ {
	// 	alphabet += string(i)
	// }
	// switch{
	// 	case c <- n:
	// 	n++
	// }

	//return string(alphabet)
	c <- n
}
func genNumbers(c chan<- string, n string) {
	// var number string
	// for i := 1; i <= 26; i++ {
	// 	number += strconv.Itoa(i)
	// }
	//return number
	c <- n
}
func PrintChannel(c <-chan string) { //<-chan  as were are reading
	for {
		msg := <-c
		fmt.Print(msg)
	}
}
func main() {
	c := make(chan string)
	num := 65
	x := 1
	for num < 91 {
		go genAlphabets(c, string(num))
		go genNumbers(c, strconv.Itoa(x))
		go PrintChannel(c)
		num++
		x++

		time.Sleep(1 * time.Millisecond)
	}

}

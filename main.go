package main

import "fmt"

func add(a int, b int) int {
	switch {
	case b > 0:
		for b > 0 {
			a++
			b--
		}

	case b < 0:
		for b < 0 {
			a--
			b++
		}
	}
	return a
}
func main() {
	a, b := 5, -10

	sum := add(a, b)
	fmt.Println("Sum:", sum)

}

package main

import (
	"fmt"
	"strconv"
)

func IntegerToBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}
func main() {
	num := 123
	bitFormat := IntegerToBinary(num)
	//fmt.Println("Binary representation:", IntegerToBinary(num), len(a))
	var bitSlice []string
	var count int
	for _, val := range bitFormat {
		bitSlice = append(bitSlice, string(val))
		if string(val) == "1" {
			count++
		}

	}
	bitSlice = append(bitSlice, strconv.Itoa(count))
	fmt.Println("val:", bitSlice)
}

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/juju/errors"
)

func main() {
	a := "100"
	b := "200"
	func1(a, b)
	func2(a, b)
}

func func1(a, b string) {
	result1, result2, err := add2AndTimes2(a, b)
	if err != nil {
		log.Fatal(1, err.Error())
	}
	fmt.Println(result1, result2)
}

func func2(a, b string) {
	result1, result2, err := add2AndTimes2(a, b)
	if err != nil {
		log.Fatal(1, errors.Trace(err).Error())
	}
	fmt.Println(result1, result2)
}

func add2AndTimes2(str1, str2 string) (int, int, error) {
	num1, err := strconv.Atoi(str1)
	num2, err := strconv.Atoi(str2)
	if err != nil {
		return 0, 0, err
	}
	return num1 + 2, num2 * 2, nil
}

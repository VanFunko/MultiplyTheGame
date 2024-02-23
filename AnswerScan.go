package main

import (
	"fmt"
	"strconv"
)

func AnswerScan(message string) int {
	var str string
	if message != "" {
		fmt.Println(message)
	}
	fmt.Scan(&str)

	intScan, err := strconv.Atoi(str)
	if err != nil {
		return AnswerScan("Вводить можно только цифры.")
	}
	return intScan
}

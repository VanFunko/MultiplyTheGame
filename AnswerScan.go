package main

import (
	"fmt"
	"os"
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
	if intScan == 0 {
		os.Exit(1)
	}
	return intScan
}

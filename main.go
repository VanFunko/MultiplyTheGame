package main

import (
	"fmt"
)

func main() {
	l := NewLogg(false)
	Clear()
	fmt.Println("В этой программе тебе нужно будет решить 5 случайно выбранных примеров из таблицы умножения.")
	StartGame(l)
}

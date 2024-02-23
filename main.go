package main

import (
	"fmt"
)

func main() {
	l := NewLogg(false)
	Clear()
	fmt.Println("Тебе нужно решить 5 случайно выбранных примеров из таблицы умножения.")
	StartGame(l)
}

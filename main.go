package main

import (
	"fmt"
)

func main() {
	l := NewLogg(false)
	fmt.Println("В этой программе тебе нужно будет решить 5 случайно выбранных примеров из таблицы умножения. Программа завершится, когда все ответы будут верны.")
	i := 5
	expMap := ChoiseQuantity(l, MultiplyTableGen(l, AnswerScan("Введи разряд таблицы или 0 для тренировки по всей таблице умножения:")), i)
	StartGame(l, expMap)
}

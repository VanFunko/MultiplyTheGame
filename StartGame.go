package main

import (
	"fmt"
)

func StartGame(l Logger) {
	counter := 1
	table := MultiplyTableGen(l, AnswerScan("Введи разряд таблицы (2-9) или 1 для тренировки по всей таблице умножения. Для выхода введи 0"))
	Clear()
	for counter > 0 { //
		wrongAnswer := make(map[int]expression)
		randomExps := ChoiceQuantity(l, table, 5) //Выбираем 5 рандомных примера
		fmt.Println("Для выхода введи 0")
		fmt.Println("Реши следующие примеры:")
		for key, _ := range randomExps {
			fmt.Printf("%v * %v = ", randomExps[key].A, randomExps[key].B)
			answer := AnswerScan("")
			if answer != randomExps[key].Res {
				wrongAnswer[key] = randomExps[key]
			} else {
				delete(wrongAnswer, key)
			}
			counter = len(wrongAnswer)
		}
		Clear()
		l.ShowLog(wrongAnswer)
		if len(wrongAnswer) > 0 {
			fmt.Printf("Количество не верных ответов: %v\n", len(wrongAnswer))
			fmt.Println("Вот верное решение:")
			for key, _ := range wrongAnswer {
				fmt.Printf("%v * %v = %v\n", wrongAnswer[key].A, wrongAnswer[key].B, wrongAnswer[key].Res)
			}
			fmt.Println("Запомни и потренируйся еще! Удачи!")
			AnswerScan("Для выхода введи 0. Чтобы продолжить введи любое число")

		} else {
			fmt.Println("Все верно! Ты молодец! УРА УРА УРА")
			rate := AnswerScan("Для выхода введи 0, чтобы продолжить изучение введи разряд табицы (2-9) или 1 для тренировки по всей таблице умножения")
			counter = 1
			clearMap(table)
			table = MultiplyTableGen(l, rate)

		}
		clearMap(wrongAnswer)
		l.ShowLog(counter)

		Clear()
	}
}

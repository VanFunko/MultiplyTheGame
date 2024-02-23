package main

import (
	"fmt"
	"os"
)

func StartGame(l Logger) {
	counter := 1
	table := MultiplyTableGen(l, AnswerScan("Введи разряд таблицы или 0 для тренировки по всей таблице умножения:"))
	for counter > 0 { //
		wrongAnswer := make(map[int]expression)
		randomExps := ChoiseQuantity(l, table, 5) //Выбираем 5 рандомных примера
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
			if AnswerScan("Для выхода введи 0. Чтобы продолжить введи любое число") == 0 {
				os.Exit(1)
			}

		} else {
			fmt.Println("Все верно! Ты молодец! УРА УРА УРА")
			if rate := AnswerScan("Для выхода введи 0, чтобы продолжить изучение введи разряд табицы"); rate != 0 {
				counter = 1
				clearMap(table)
				table = MultiplyTableGen(l, rate)
			}
		}
		clearMap(wrongAnswer)
		l.ShowLog(counter)

		Clear()
	}
}

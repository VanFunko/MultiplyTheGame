package main

import "fmt"

func StartGame(l Logger, expMap map[int]expression) {
	counter := 1
	wrongAnswer := make(map[int]expression)
	for counter > 0 {
		fmt.Println("Реши следующие примеры")
		for key, _ := range expMap {
			fmt.Printf("%v * %v = ", expMap[key].A, expMap[key].B)
			answer := AnswerScan("")
			if answer != expMap[key].Res {
				wrongAnswer[key] = expMap[key]
			} else {
				delete(wrongAnswer, key)
			}
			counter = len(wrongAnswer)
		}
		l.ShowLog(wrongAnswer)
		if len(wrongAnswer) > 0 {
			fmt.Println("_________________________________________________________________________")
			fmt.Printf("Количество не верных ответов: %v\n", len(wrongAnswer))
			fmt.Println("Вот верное решение:")
			for key, _ := range wrongAnswer {
				fmt.Printf("%v * %v = %v\n", wrongAnswer[key].A, wrongAnswer[key].B, wrongAnswer[key].Res)
			}
			fmt.Println("Запомни и потренируйся еще! Удачи!")
			fmt.Println("")
			fmt.Println("_________________________________________________________________________")
		} else {
			fmt.Println("Все верно! Ты молодец! УРА УРА УРА")
		}
		l.ShowLog(counter)
	}
}

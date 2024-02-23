package main

func clearMap(m map[int]expression) map[int]expression {
	for k := range m { // Обнуляем карту не правильных ответов
		delete(m, k)
	}
	return m
}

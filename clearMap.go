package main

func clearMap(m map[int]expression) map[int]expression {
	for k := range m {
		delete(m, k)
	}
	return m
}

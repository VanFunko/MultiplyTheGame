package main

import "math/rand"

func ChoiseQuantity(l Logger, allExps map[int]expression, quantityExps int) map[int]expression {
	expMap := make(map[int]expression)
	uniqMap := make(map[int]int)
	for len(uniqMap) < quantityExps {
		id := rand.Intn(len(allExps)-1) + 1
		uniqMap[id]++
	}
	l.ShowLog(uniqMap)
	for key, _ := range uniqMap {
		expMap[key] = allExps[key]
	}
	//for i := 1; i <= quantityExps; i++ {
	//	id := rand.Intn(len(allExps)-1) + 1
	//	expMap[i] = allExps[id]
	//}
	l.ShowLog(expMap)
	return expMap
}

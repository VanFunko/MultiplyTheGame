package main

type expression struct {
	A   int `json:"a"`
	B   int `json:"b"`
	Res int `json:"res"`
}

func MultiplyTableGen(logg Logger, from int) map[int]expression {
	expMap := make(map[int]expression)
	exp := expression{}
	ID := 0
	if from == 1 {
		for i := 2; i <= 10; i++ {
			for j := 2; j <= 10; j++ {
				ID++
				exp.A = i
				exp.B = j
				exp.Res = exp.A * exp.B
				expMap[ID] = exp
			}
		}
	} else {
		for j := 2; j <= 10; j++ {
			ID++
			exp.A = from
			exp.B = j
			exp.Res = exp.A * exp.B
			expMap[ID] = exp
		}
	}
	logg.ShowLog(expMap)
	return expMap
}

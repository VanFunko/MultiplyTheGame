package main

import (
	"fmt"
	"log"
)

type Logger interface {
	ShowLog(message any)
}

type Logg struct {
	isShow bool
}

func NewLogg(isShow bool) *Logg {
	return &Logg{isShow: isShow}
}

func (l *Logg) ShowLog(logMessages any) {
	if l.isShow {
		message := fmt.Sprintf("\x1b[33m%v\x1b[0m ", logMessages)
		log.Println(message)
	}

}

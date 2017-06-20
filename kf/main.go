package main

import (
	"fca/dal"
	"fca/libs"
	"fca/logic"
	"time"
)

var Logic logic.LogicBase

func KFRoutine() {
	for {
		Logic.KF()
		time.Sleep(1 * time.Second)
	}
}
func main() {

	libs.NewLogger()
	Logic = logic.LogicBase{}
	dal.InitDb()
	logic.InitLogic()
	go KFRoutine()

	c := make(chan int)
	<-c
}

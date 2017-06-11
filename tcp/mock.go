package main

import (
	"fmt"
	"time"
)

var t1 chan int = make(chan int, 100)
var t2 chan int = make(chan int, 100)

type TestPole struct {
	PoleID string
	Cmd    int
}

func Mock9(uuid string, t chan int) {
	time1 := 0
	b := true
	for b {
		time1++
		select {
		case <-t:
			fmt.Printf("stop:%s\n", uuid)
			b = false
			break
		default:
			time.Sleep(100 * time.Microsecond)
			Logic.UpdateCharge(uuid, "1", int32(time1), float32(time1), int32(time1))
		}
	}

}

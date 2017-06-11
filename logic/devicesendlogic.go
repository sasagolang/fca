package logic

import (
	"fca/libs"
	"fca/model/proto"
	"time"

	"fmt"

	protobuf "github.com/golang/protobuf/proto"
)

func (logic LogicBase) Start_cmd(mins int32, time time.Time, ch chan *libs.Proto) {
	start1 := proto.START{}
	start1.Mins = protobuf.Int32(mins)
	start1.Time = time2TIME(time)
	p := &libs.Proto{}
	p.CMD = 4
	p.Content, _ = protobuf.Marshal(&start1)
	ch <- p

}
func (logic LogicBase) Stop_cmd(ch chan *libs.Proto) {
	p := &libs.Proto{}
	p.CMD = 5
	p.Content = make([]byte, 0)
	ch <- p
}
func (logic LogicBase) DownConfig_cmd(config *proto.CONFIG, ch chan *libs.Proto) {

	p := &libs.Proto{}
	p.CMD = 8
	p.Content, _ = protobuf.Marshal(config)
	ch <- p

}
func (logic LogicBase) StartLight_cmd(ch chan *libs.Proto) {
	fmt.Printf("StartLight_cmd:%v\n", ch)
	p := &libs.Proto{}
	p.CMD = 6
	p.Content = make([]byte, 0)
	ch <- p
}
func (logic LogicBase) StopLight_cmd(ch chan *libs.Proto) {
	p := &libs.Proto{}
	p.CMD = 7
	p.Content = make([]byte, 0)
	ch <- p
}

package logic

import (
	"fca/libs"
	"fca/model/proto"
	"fmt"
	"time"

	protobuf "github.com/golang/protobuf/proto"
)

func SendMsg(uuid string, p *libs.Proto) {

}
func Start_cmd(mins int32, time time.Time, ch chan *libs.Proto) {
	start1 := proto.START{}
	start1.Mins = protobuf.Int32(30)
	start1.Time = time2TIME(time)
	p := &libs.Proto{}
	p.CMD = 4
	p.Content, _ = protobuf.Marshal(&start1)
	ch <- p

}
func ReceivedMsg(p *libs.Proto, ch chan *libs.Proto) {
	var data interface{}
	var err error
	//fmt.Printf("ReceivedMsg:%v\n", p)
	switch p.CMD {
	case 1: //上行握手
		data, err = Handshake_up(p, ch)
		fmt.Printf("ReceivedMsg:CMD:%v,err:%v,data:%v\n", p.CMD, err, data.(proto.HANDSHAKE_UP))
	case 2: //心跳
		HeartBeat(p, ch)
	case 3: //设备状态数据
		data, err = Base(p, ch)
		Start_cmd(1, time.Now().Add(10*time.Minute), ch)
	case 8: // 设备属性配置数据
		data, err = Confi_up(p, ch)
	case 9: // 充电中实时数据
		data, err = Charge(p, ch)
	case 10: //设备告警数据
		data, err = Alarm(p, ch)
	case 11: //电表实时数据
		data, err = Meter(p, ch)
	case 12: //电桩交易数据
		data, err = Trade(p, ch)
	case 13: //设备固件升级
		data, err = Upgrade_down(p, ch)

	}
	fmt.Printf("ReceivedMsg:CMD:%v,err:%v,data:%v\n", p.CMD, err, &data)

}
func time2TIME(time time.Time) *proto.TIME {
	t := proto.TIME{}
	t.Year = protobuf.Int32(int32(time.Year()))
	t.Month = protobuf.Int32(int32(time.Month()))

	t.Day = protobuf.Int32(int32(time.Day()))
	t.Hour = protobuf.Int32(int32(time.Hour()))
	t.Min = protobuf.Int32(int32(time.Minute()))
	t.Sec = protobuf.Int32(int32(time.Second()))
	return &t
}
func TIME2Time(t proto.TIME) time.Time {
	return time.Date(int(*t.Year), time.Month(*t.Month), int(*t.Day), int(*t.Hour), int(*t.Min), int(*t.Sec), 0, time.Local)

}
func Upgrade_down(p *libs.Proto, ch chan *libs.Proto) (pb proto.UPGRADE_DOWN, err error) {
	err = protobuf.Unmarshal(p.Content, &pb)
	return
}
func Trade(p *libs.Proto, ch chan *libs.Proto) (pb proto.TRADE, err error) {
	err = protobuf.Unmarshal(p.Content, &pb)
	return
}
func Meter(p *libs.Proto, ch chan *libs.Proto) (pb proto.METER, err error) {
	err = protobuf.Unmarshal(p.Content, &pb)
	return
}
func Alarm(p *libs.Proto, ch chan *libs.Proto) (pb proto.ALARM, err error) {
	err = protobuf.Unmarshal(p.Content, &pb)
	return
}
func Base(p *libs.Proto, ch chan *libs.Proto) (pb proto.BASE, err error) {
	err = protobuf.Unmarshal(p.Content, &pb)
	return
}
func Confi_up(p *libs.Proto, ch chan *libs.Proto) (pb proto.CONFIG, err error) {
	err = protobuf.Unmarshal(p.Content, &pb)
	return
}
func Charge(p *libs.Proto, ch chan *libs.Proto) (pb proto.CHARGE, err error) {
	err = protobuf.Unmarshal(p.Content, &pb)
	return
}
func HeartBeat(p *libs.Proto, ch chan *libs.Proto) {
	po := &libs.Proto{}
	po.CMD = 2
	po.Content = make([]byte, 0)
	po.UDID = p.UDID

	ch <- po
}
func Handshake_up(p *libs.Proto, ch chan *libs.Proto) (pb proto.HANDSHAKE_UP, err error) {

	err = protobuf.Unmarshal(p.Content, &pb)
	//发送下行握手
	s := proto.HANDSHAKE_DOWN{}
	s.Ip = protobuf.String("192.168.1.146")
	s.Port = protobuf.Int32(8998)
	s.Time = time2TIME(time.Now())
	b, er := protobuf.Marshal(&s)
	if er != nil {
		panic(er)
	}
	po := &libs.Proto{}
	po.CMD = 1
	po.Content = b
	po.UDID = p.UDID
	ch <- po
	return
}

func Handshake_down(buf []byte) (pb proto.HANDSHAKE_DOWN, err error) {
	err = protobuf.Unmarshal(buf, &pb)
	return
}

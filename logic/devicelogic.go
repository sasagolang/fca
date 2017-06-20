package logic

import (
	"fca/libs"
	"fca/model/proto"
	"fmt"
	"time"

	"strconv"

	protobuf "github.com/golang/protobuf/proto"
)

func (logic LogicBase) ReceivedMsg(p *libs.Proto, ch chan *libs.Proto) {
	var data interface{}
	var err error
	//fmt.Printf("ReceivedMsg:%v\n", p)
	switch p.CMD {
	case 1: //上行握手
		data, err = Handshake_up(p, ch)
	//	fmt.Printf("ReceivedMsg:CMD:%v,err:%v,data:%v\n", p.CMD, err, data.(proto.HANDSHAKE_UP))
	case 2: //心跳
		HeartBeat(p, ch)
	case 3: //设备状态数据
		data, err = Base(p, ch)
		if err == nil {
			v := data.(proto.BASE)
			//更新电桩信息
			var statusname string
			switch v.GetMode() {
			case 0:
				statusname = "异常"

			case 1:
				statusname = "空闲"

			case 2:
				statusname = "充电中"

			case 3:
				statusname = "固件升级中"

			case 4:
				statusname = "充电枪已连接"

			case 5:
				statusname = "充满"

			case 6:
				statusname = "充电结束"
				//停止充电订单
				logic.UpdateCharge(p.UDID, "102", 0, 0, 0)
			default:
				statusname = "未知"
			}
			logic.UpdatePoleStatus(p.UDID, int32(v.GetMode()+2000), statusname)

		}
		//Start_cmd(1, time.Now().Add(10*time.Minute), ch)
	case 8: // 设备属性配置数据
		data, err = Confi_up(p, ch)
	case 9: // 充电中实时数据
		data, err = Charge(p, ch)
		v := data.(proto.CHARGE)
		fmt.Printf("proto.CHARGE:%v\n", v)
		logic.UpdateCharge(p.UDID, strconv.Itoa(int(*v.State)), *v.Energy, *v.Amount, *v.Duration)
	case 10: //设备告警数据
		data, err = Alarm(p, ch)
	case 11: //电表实时数据
		data, err = Meter(p, ch)
	case 12: //电桩交易数据
		data, err = Trade(p, ch)
	case 13: //设备固件升级
		data, err = Upgrade_down(p, ch)

	}
	if err != nil {
		fmt.Printf("ReceivedMsg:CMD:%v,err:%v,data:%v\n", p.CMD, err, &data)
	}

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
	po := &libs.Proto{}
	po.CMD = 13
	po.Content = make([]byte, 0)
	po.UDID = p.UDID

	ch <- po
	return
}
func Trade(p *libs.Proto, ch chan *libs.Proto) (pb proto.TRADE, err error) {
	err = protobuf.Unmarshal(p.Content, &pb)
	fmt.Printf("Trade:%v\n", &pb)
	po := &libs.Proto{}
	po.CMD = 12
	po.Content = make([]byte, 0)
	po.UDID = p.UDID

	ch <- po
	return
}
func Meter(p *libs.Proto, ch chan *libs.Proto) (pb proto.METER, err error) {
	err = protobuf.Unmarshal(p.Content, &pb)
	po := &libs.Proto{}
	po.CMD = 11
	po.Content = make([]byte, 0)
	po.UDID = p.UDID

	ch <- po
	return
}
func Alarm(p *libs.Proto, ch chan *libs.Proto) (pb proto.ALARM, err error) {
	err = protobuf.Unmarshal(p.Content, &pb)
	po := &libs.Proto{}
	po.CMD = 10
	po.Content = make([]byte, 0)
	po.UDID = p.UDID

	ch <- po
	return
}
func Base(p *libs.Proto, ch chan *libs.Proto) (pb proto.BASE, err error) {
	err = protobuf.Unmarshal(p.Content, &pb)

	po := &libs.Proto{}
	po.CMD = 3
	po.Content = make([]byte, 0)
	po.UDID = p.UDID

	ch <- po
	return
}
func Confi_up(p *libs.Proto, ch chan *libs.Proto) (pb proto.CONFIG, err error) {
	err = protobuf.Unmarshal(p.Content, &pb)
	po := &libs.Proto{}
	fmt.Printf("Confi_up:%v\n", &pb)
	po.CMD = 8
	po.Content = make([]byte, 0)
	po.UDID = p.UDID

	//ch <- po
	return
}
func Charge(p *libs.Proto, ch chan *libs.Proto) (pb proto.CHARGE, err error) {
	err = protobuf.Unmarshal(p.Content, &pb)
	po := &libs.Proto{}
	po.CMD = 9
	po.Content = make([]byte, 0)
	po.UDID = p.UDID

	ch <- po
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
	s.Ip = protobuf.String("119.23.22.219")
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

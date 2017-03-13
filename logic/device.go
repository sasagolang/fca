package logic

import (
	"encoding/binary"
	"fca/libs"
	"fca/model/proto"
	"fmt"
	"time"
	"unsafe"

	"bytes"

	protobuf "github.com/golang/protobuf/proto"
)

type DeviceProto struct {
	Identify  byte
	Ver       byte
	PackeLen  int32
	TimeStamp uint32
	UDType    uint16
	UDID      string
	Content   []byte
	CMD       uint16
	CRC       uint16
}

func (d *DeviceProto) Handshake_up(imei string, imsi string, report string) ([]byte, error) {

	pb := proto.HANDSHAKE_UP{}
	pb.Imei = protobuf.String(imei)
	pb.Imsi = protobuf.String(imsi)
	pb.Report = protobuf.String(report)
	return protobuf.Marshal(&pb)
}
func (d *DeviceProto) Handshake_down(port int32, ip string, time time.Time) ([]byte, error) {
	t := proto.TIME{}
	t.Year = protobuf.Int32(int32(time.Year()))
	t.Month = protobuf.Int32(int32(time.Month()))

	t.Day = protobuf.Int32(int32(time.Day()))
	t.Hour = protobuf.Int32(int32(time.Hour()))
	t.Min = protobuf.Int32(int32(time.Minute()))
	t.Sec = protobuf.Int32(int32(time.Second()))
	pb := proto.HANDSHAKE_DOWN{}
	pb.Port = protobuf.Int32(port)
	pb.Ip = protobuf.String(ip)
	pb.Time = &t
	return protobuf.Marshal(&pb)

}

func (d *DeviceProto) Parse() (data interface{}, err error) {
	if d.CMD == 1 {
		var data proto.HANDSHAKE_UP
		protobuf.Unmarshal(d.Content, &data)
		fmt.Printf("ContentParse:%v", &data)
		libs.Log.Infof("ContentParse:%v", &data)
		return &data, nil
	}
	return nil, nil
}

func (d *DeviceProto) start_cmd(mins string, time time.Time) {
	//pb := proto.START{}
}

var SX byte = 'K'
var XX byte = 'H'
var Ver byte = byte(1)
var RawHeaderSize int = 32

func (d *DeviceProto) BuildHead(udtype uint16, udid string, cmd uint16, content []byte) (buff []byte) {
	b1 := new(bytes.Buffer)
	//上行
	//buf[0] = SX
	binary.Write(b1, binary.BigEndian, SX)
	//协议版本
	//buf[1] = Ver
	binary.Write(b1, binary.BigEndian, Ver)
	//包长
	pl := len(content) + RawHeaderSize
	libs.Log.Error(fmt.Sprintf("BuildHead.pl:%v\n", pl))
	binary.Write(b1, binary.BigEndian, int32(pl))
	//时间戳
	binary.Write(b1, binary.BigEndian, int32(time.Now().Nanosecond()))
	//电桩类型
	binary.Write(b1, binary.BigEndian, udtype)
	//设备码
	imei := udid
	for index := len(imei); index < 16; index++ {
		imei = "0" + imei
	}

	tmp := []byte(imei)
	libs.Log.Info(fmt.Sprintf("BuildHead.tmp:%d,%v\n", len(imei), tmp))
	b1.Write(tmp)
	//	binary.Write(b1, binary.BigEndian, imei)
	//命令字
	binary.Write(b1, binary.BigEndian, cmd)
	//报文数据内容
	binary.Write(b1, binary.BigEndian, content)
	//校验数据
	crc := libs.Crc16(b1.Bytes())
	binary.Write(b1, binary.BigEndian, crc)
	libs.Log.Error(fmt.Sprintf("BuildHead%v\n", b1.Bytes()))
	return b1.Bytes()

}
func Int32(b []byte) int32 {
	return int32(b[3]) | int32(b[2])<<8 | int32(b[1])<<16 | int32(b[0])<<24
}
func BytesString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func (d *DeviceProto) ParseHead(buff []byte) {

	identify := buff[0]

	ver := buff[1]
	if identify != XX {

	}
	if ver != Ver {

	}
	pl := binary.BigEndian.Uint32(buff[2:6])

	timestamp := binary.BigEndian.Uint32(buff[6:10])
	udtype := binary.BigEndian.Uint16(buff[10:12])
	udid := string(buff[12:28])
	cmd := binary.BigEndian.Uint16(buff[28:30])
	content := buff[30 : int(pl)-RawHeaderSize+30]
	crc := binary.BigEndian.Uint16(buff[int(pl)-RawHeaderSize+30 : pl])
	d.Identify = identify
	d.Ver = ver
	d.PackeLen = int32(pl)
	d.TimeStamp = timestamp
	d.UDType = udtype
	d.UDID = udid
	d.CMD = cmd
	d.Content = content
	d.CRC = crc

}

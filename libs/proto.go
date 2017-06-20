package libs

import (
	bu "bytes"
	"encoding/binary"
	"fca/libs/bufio"
	"fmt"
	"strings"
	"time"
)

type Proto struct {
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

func (p *Proto) ReadTCP(rr *bufio.Reader) (re *Proto, err error) {
	re = &Proto{}
	var buf []byte

	if buf, err = rr.Pop(30); err != nil {

		return
	}

	//	fmt.Printf("ReadTCP%v,%v\n", p, buf)
	re.Identify = buf[0]
	re.Ver = buf[1]
	re.PackeLen = int32(binary.BigEndian.Uint32(buf[2:6]))
	//binary.Read(rr, binary.BigEndian, p.Identify)
	//binary.Read(rr, binary.BigEndian, p.Ver)
	//binary.Read(rr, binary.BigEndian, pl)
	re.TimeStamp = binary.BigEndian.Uint32(buf[6:10])
	re.UDType = binary.BigEndian.Uint16(buf[10:12])
	//binary.Read(rr, binary.BigEndian, p.TimeStamp)
	//binary.Read(rr, binary.BigEndian, p.UDType)
	//fmt.Printf("re.UDID:%v\n", buf[12:28])
	re.UDID = string(buf[12:28])
	re.UDID = strings.Replace(re.UDID, string([]byte{0}), "", -1)
	re.CMD = binary.BigEndian.Uint16(buf[28:30])
	//udidb := [16]byte{}
	//binary.Read(rr, binary.BigEndian, udidb)
	//	p.UDID = string(udidb[0:16])
	//	binary.Read(rr, binary.BigEndian, p.CMD)

	//bytes, err := rr.Peek(int(p.PackeLen) - 32)
	if re.PackeLen > 32 {
		re.Content, err = rr.Pop(int(re.PackeLen - 32))
	}
	var crcb []byte
	if crcb, err = rr.Pop(2); err != nil {
		return
	}
	if err != nil {
		return
	}
	re.CRC = binary.BigEndian.Uint16(crcb)
	//	fmt.Printf("ReadTCP.p:%v\n", re)
	//binary.Read(rr, binary.BigEndian, p.CRC)
	return
}

var XX byte = 'H'
var Ver byte = byte(1)
var RawHeaderSize int = 32

func (p *Proto) WriteTo(b *bufio.Writer) {

	b1 := new(bu.Buffer)
	binary.Write(b1, binary.BigEndian, XX)
	binary.Write(b1, binary.BigEndian, Ver)
	pl := len(p.Content) + RawHeaderSize

	binary.Write(b1, binary.BigEndian, int32(pl))
	binary.Write(b1, binary.BigEndian, int32(time.Now().Unix()))

	binary.Write(b1, binary.BigEndian, p.UDType)
	imei := p.UDID
	for index := len(imei); index < 16; index++ {
		imei = "0" + imei
	}
	tmp := []byte(imei)
	b1.Write(tmp)
	binary.Write(b1, binary.BigEndian, p.CMD)
	binary.Write(b1, binary.BigEndian, p.Content)
	crc := Crc16(b1.Bytes())
	binary.Write(b1, binary.BigEndian, crc)
	_, err := b.Write(b1.Bytes())
	fmt.Printf("WriteTo:%v\n", p.CMD)
	if err != nil {
		fmt.Printf("WriteTo.error:%v\n", err)
	}
}

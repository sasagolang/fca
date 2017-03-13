package logic

import (
	"fca/model/proto"
	"testing"

	"time"

	protobuf "github.com/golang/protobuf/proto"
)

var devicet *DeviceProto = &DeviceProto{}

func Test_Handshake_up(t *testing.T) {
	bytes, _ := devicet.Handshake_up("123", "456", "127.0.0.1")
	var pbn proto.HANDSHAKE_UP
	protobuf.Unmarshal(bytes, &pbn)
	if *pbn.Imei != "123" || *pbn.Imsi != "456" || *pbn.Report != "127.0.0.1" {
		t.Error("Handshake_up 测试失败")
	}

}
func Test_Handshake_down(t *testing.T) {
	//port int32, ip string, time time.Time
	port := int32(123)
	ip := "127.0.0.1"
	tt := time.Date(2017, 3, 4, 13, 20, 21, 0, time.Local)
	bytes, _ := devicet.Handshake_down(port, ip, tt)
	var pbn proto.HANDSHAKE_DOWN
	protobuf.Unmarshal(bytes, &pbn)
	t.Log(&tt)
	t.Log(pbn.Time)
	if *pbn.Ip != ip || *pbn.Port != port || *pbn.Time.Year != int32(tt.Year()) {
		t.Error("Handshake_down 测试失败")
	}
}

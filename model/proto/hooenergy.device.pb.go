// Code generated by protoc-gen-go.
// source: hooenergy.device.proto
// DO NOT EDIT!

/*
Package hooenergy_device is a generated protocol buffer package.

It is generated from these files:
	hooenergy.device.proto

It has these top-level messages:
	TIME
	HANDSHAKE_UP
	HANDSHAKE_DOWN
	UPGRADE_UP
	UPGRADE_DOWN
	START
	BASE
	FEERATE
	PERIOD
	CONFIG
	CHARGE
	ALARM
	METER
	TRADE
*/
package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BASE_MODE int32

const (
	BASE_STATUS_EXCEPTION BASE_MODE = 0
	BASE_STATUS_IDLE      BASE_MODE = 1
	BASE_STATUS_CHARGE    BASE_MODE = 2
	BASE_STATUS_UPGRADE   BASE_MODE = 3
	BASE_STATUS_CONNECTED BASE_MODE = 4
	BASE_STATUS_COMPLETE  BASE_MODE = 5
	BASE_STATUS_FINISH    BASE_MODE = 6
)

var BASE_MODE_name = map[int32]string{
	0: "STATUS_EXCEPTION",
	1: "STATUS_IDLE",
	2: "STATUS_CHARGE",
	3: "STATUS_UPGRADE",
	4: "STATUS_CONNECTED",
	5: "STATUS_COMPLETE",
	6: "STATUS_FINISH",
}
var BASE_MODE_value = map[string]int32{
	"STATUS_EXCEPTION": 0,
	"STATUS_IDLE":      1,
	"STATUS_CHARGE":    2,
	"STATUS_UPGRADE":   3,
	"STATUS_CONNECTED": 4,
	"STATUS_COMPLETE":  5,
	"STATUS_FINISH":    6,
}

func (x BASE_MODE) Enum() *BASE_MODE {
	p := new(BASE_MODE)
	*p = x
	return p
}
func (x BASE_MODE) String() string {
	return proto.EnumName(BASE_MODE_name, int32(x))
}
func (x *BASE_MODE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BASE_MODE_value, data, "BASE_MODE")
	if err != nil {
		return err
	}
	*x = BASE_MODE(value)
	return nil
}
func (BASE_MODE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

type BASE_STATUS int32

const (
	BASE_STATUS_UNKNOWN       BASE_STATUS = 0
	BASE_STATUS_DISCONNUNLOCK BASE_STATUS = 1
	BASE_STATUS_CONNUNLOCK    BASE_STATUS = 2
	BASE_STATUS_CONNLOCK      BASE_STATUS = 3
	BASE_STATUS_DISCONNLOCK   BASE_STATUS = 4
)

var BASE_STATUS_name = map[int32]string{
	0: "STATUS_UNKNOWN",
	1: "STATUS_DISCONNUNLOCK",
	2: "STATUS_CONNUNLOCK",
	3: "STATUS_CONNLOCK",
	4: "STATUS_DISCONNLOCK",
}
var BASE_STATUS_value = map[string]int32{
	"STATUS_UNKNOWN":       0,
	"STATUS_DISCONNUNLOCK": 1,
	"STATUS_CONNUNLOCK":    2,
	"STATUS_CONNLOCK":      3,
	"STATUS_DISCONNLOCK":   4,
}

func (x BASE_STATUS) Enum() *BASE_STATUS {
	p := new(BASE_STATUS)
	*p = x
	return p
}
func (x BASE_STATUS) String() string {
	return proto.EnumName(BASE_STATUS_name, int32(x))
}
func (x *BASE_STATUS) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BASE_STATUS_value, data, "BASE_STATUS")
	if err != nil {
		return err
	}
	*x = BASE_STATUS(value)
	return nil
}
func (BASE_STATUS) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 1} }

type CONFIG_CATEGORY int32

const (
	CONFIG_DEFAULT_UNKNOWN CONFIG_CATEGORY = 0
	CONFIG_HANGING_AC      CONFIG_CATEGORY = 1
)

var CONFIG_CATEGORY_name = map[int32]string{
	0: "DEFAULT_UNKNOWN",
	1: "HANGING_AC",
}
var CONFIG_CATEGORY_value = map[string]int32{
	"DEFAULT_UNKNOWN": 0,
	"HANGING_AC":      1,
}

func (x CONFIG_CATEGORY) Enum() *CONFIG_CATEGORY {
	p := new(CONFIG_CATEGORY)
	*p = x
	return p
}
func (x CONFIG_CATEGORY) String() string {
	return proto.EnumName(CONFIG_CATEGORY_name, int32(x))
}
func (x *CONFIG_CATEGORY) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CONFIG_CATEGORY_value, data, "CONFIG_CATEGORY")
	if err != nil {
		return err
	}
	*x = CONFIG_CATEGORY(value)
	return nil
}
func (CONFIG_CATEGORY) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{9, 0} }

type CHARGE_STATE int32

const (
	CHARGE_CHARGING_EXCEPTION CHARGE_STATE = 0
	CHARGE_CHARGING_ON        CHARGE_STATE = 1
	CHARGE_CHARGING_COMPLETE  CHARGE_STATE = 2
)

var CHARGE_STATE_name = map[int32]string{
	0: "CHARGING_EXCEPTION",
	1: "CHARGING_ON",
	2: "CHARGING_COMPLETE",
}
var CHARGE_STATE_value = map[string]int32{
	"CHARGING_EXCEPTION": 0,
	"CHARGING_ON":        1,
	"CHARGING_COMPLETE":  2,
}

func (x CHARGE_STATE) Enum() *CHARGE_STATE {
	p := new(CHARGE_STATE)
	*p = x
	return p
}
func (x CHARGE_STATE) String() string {
	return proto.EnumName(CHARGE_STATE_name, int32(x))
}
func (x *CHARGE_STATE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CHARGE_STATE_value, data, "CHARGE_STATE")
	if err != nil {
		return err
	}
	*x = CHARGE_STATE(value)
	return nil
}
func (CHARGE_STATE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{10, 0} }

type ALARM_TYPE int32

const (
	ALARM_NETWORK_FAULT     ALARM_TYPE = 1
	ALARM_SCREEN_FAULT      ALARM_TYPE = 2
	ALARM_METER_FAULT       ALARM_TYPE = 3
	ALARM_INPUT_OVER        ALARM_TYPE = 4
	ALARM_INPUT_UNDER       ALARM_TYPE = 5
	ALARM_OUTPUT_OVER       ALARM_TYPE = 6
	ALARM_OUTPUT_UNDER      ALARM_TYPE = 7
	ALARM_OUTPUT_OVER_CUR   ALARM_TYPE = 8
	ALARM_SPD_FAULT         ALARM_TYPE = 9
	ALARM_TEMP_EXCEED       ALARM_TYPE = 10
	ALARM_TEMP_SENSOR_FAULT ALARM_TYPE = 11
	ALARM_TEMP_LOW          ALARM_TYPE = 12
	ALARM_TEMP_SPEAR_EXCEED ALARM_TYPE = 13
	ALARM_CRASH_STOP        ALARM_TYPE = 14
	ALARM_CARD_FAULT        ALARM_TYPE = 15
	ALARM_LOCK_FAULT        ALARM_TYPE = 16
)

var ALARM_TYPE_name = map[int32]string{
	1:  "NETWORK_FAULT",
	2:  "SCREEN_FAULT",
	3:  "METER_FAULT",
	4:  "INPUT_OVER",
	5:  "INPUT_UNDER",
	6:  "OUTPUT_OVER",
	7:  "OUTPUT_UNDER",
	8:  "OUTPUT_OVER_CUR",
	9:  "SPD_FAULT",
	10: "TEMP_EXCEED",
	11: "TEMP_SENSOR_FAULT",
	12: "TEMP_LOW",
	13: "TEMP_SPEAR_EXCEED",
	14: "CRASH_STOP",
	15: "CARD_FAULT",
	16: "LOCK_FAULT",
}
var ALARM_TYPE_value = map[string]int32{
	"NETWORK_FAULT":     1,
	"SCREEN_FAULT":      2,
	"METER_FAULT":       3,
	"INPUT_OVER":        4,
	"INPUT_UNDER":       5,
	"OUTPUT_OVER":       6,
	"OUTPUT_UNDER":      7,
	"OUTPUT_OVER_CUR":   8,
	"SPD_FAULT":         9,
	"TEMP_EXCEED":       10,
	"TEMP_SENSOR_FAULT": 11,
	"TEMP_LOW":          12,
	"TEMP_SPEAR_EXCEED": 13,
	"CRASH_STOP":        14,
	"CARD_FAULT":        15,
	"LOCK_FAULT":        16,
}

func (x ALARM_TYPE) Enum() *ALARM_TYPE {
	p := new(ALARM_TYPE)
	*p = x
	return p
}
func (x ALARM_TYPE) String() string {
	return proto.EnumName(ALARM_TYPE_name, int32(x))
}
func (x *ALARM_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ALARM_TYPE_value, data, "ALARM_TYPE")
	if err != nil {
		return err
	}
	*x = ALARM_TYPE(value)
	return nil
}
func (ALARM_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{11, 0} }

type TRADE_TYPE int32

const (
	TRADE_USER_APP TRADE_TYPE = 1
	TRADE_USER_IC  TRADE_TYPE = 2
)

var TRADE_TYPE_name = map[int32]string{
	1: "USER_APP",
	2: "USER_IC",
}
var TRADE_TYPE_value = map[string]int32{
	"USER_APP": 1,
	"USER_IC":  2,
}

func (x TRADE_TYPE) Enum() *TRADE_TYPE {
	p := new(TRADE_TYPE)
	*p = x
	return p
}
func (x TRADE_TYPE) String() string {
	return proto.EnumName(TRADE_TYPE_name, int32(x))
}
func (x *TRADE_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TRADE_TYPE_value, data, "TRADE_TYPE")
	if err != nil {
		return err
	}
	*x = TRADE_TYPE(value)
	return nil
}
func (TRADE_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{13, 0} }

type TRADE_CAUSE int32

const (
	TRADE_USER_ACTIVE       TRADE_CAUSE = 1
	TRADE_CHARGE_COMPLETE   TRADE_CAUSE = 2
	TRADE_REMAIN_NOT_ENOUGH TRADE_CAUSE = 3
	TRADE_EXCEPTION_TERM    TRADE_CAUSE = 4
	TRADE_URGENCY_TERM      TRADE_CAUSE = 5
	TRADE_PAEK_SWITCH       TRADE_CAUSE = 6
	TRADE_DATE_SWITCH       TRADE_CAUSE = 7
)

var TRADE_CAUSE_name = map[int32]string{
	1: "USER_ACTIVE",
	2: "CHARGE_COMPLETE",
	3: "REMAIN_NOT_ENOUGH",
	4: "EXCEPTION_TERM",
	5: "URGENCY_TERM",
	6: "PAEK_SWITCH",
	7: "DATE_SWITCH",
}
var TRADE_CAUSE_value = map[string]int32{
	"USER_ACTIVE":       1,
	"CHARGE_COMPLETE":   2,
	"REMAIN_NOT_ENOUGH": 3,
	"EXCEPTION_TERM":    4,
	"URGENCY_TERM":      5,
	"PAEK_SWITCH":       6,
	"DATE_SWITCH":       7,
}

func (x TRADE_CAUSE) Enum() *TRADE_CAUSE {
	p := new(TRADE_CAUSE)
	*p = x
	return p
}
func (x TRADE_CAUSE) String() string {
	return proto.EnumName(TRADE_CAUSE_name, int32(x))
}
func (x *TRADE_CAUSE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TRADE_CAUSE_value, data, "TRADE_CAUSE")
	if err != nil {
		return err
	}
	*x = TRADE_CAUSE(value)
	return nil
}
func (TRADE_CAUSE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{13, 1} }

type TIME struct {
	Year             *int32  `protobuf:"varint,1,req,name=year" json:"year,omitempty"`
	Month            *int32  `protobuf:"varint,2,req,name=month" json:"month,omitempty"`
	Day              *int32  `protobuf:"varint,3,req,name=day" json:"day,omitempty"`
	Hour             *int32  `protobuf:"varint,4,req,name=hour" json:"hour,omitempty"`
	Min              *int32  `protobuf:"varint,5,req,name=min" json:"min,omitempty"`
	Sec              *int32  `protobuf:"varint,6,req,name=sec" json:"sec,omitempty"`
	Timestamp        *uint32 `protobuf:"varint,7,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TIME) Reset()                    { *m = TIME{} }
func (m *TIME) String() string            { return proto.CompactTextString(m) }
func (*TIME) ProtoMessage()               {}
func (*TIME) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TIME) GetYear() int32 {
	if m != nil && m.Year != nil {
		return *m.Year
	}
	return 0
}

func (m *TIME) GetMonth() int32 {
	if m != nil && m.Month != nil {
		return *m.Month
	}
	return 0
}

func (m *TIME) GetDay() int32 {
	if m != nil && m.Day != nil {
		return *m.Day
	}
	return 0
}

func (m *TIME) GetHour() int32 {
	if m != nil && m.Hour != nil {
		return *m.Hour
	}
	return 0
}

func (m *TIME) GetMin() int32 {
	if m != nil && m.Min != nil {
		return *m.Min
	}
	return 0
}

func (m *TIME) GetSec() int32 {
	if m != nil && m.Sec != nil {
		return *m.Sec
	}
	return 0
}

func (m *TIME) GetTimestamp() uint32 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

// 与服务器握手数据(上行)
type HANDSHAKE_UP struct {
	Imei             *string `protobuf:"bytes,1,req,name=imei" json:"imei,omitempty"`
	Imsi             *string `protobuf:"bytes,2,req,name=imsi" json:"imsi,omitempty"`
	Report           *string `protobuf:"bytes,3,opt,name=report" json:"report,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HANDSHAKE_UP) Reset()                    { *m = HANDSHAKE_UP{} }
func (m *HANDSHAKE_UP) String() string            { return proto.CompactTextString(m) }
func (*HANDSHAKE_UP) ProtoMessage()               {}
func (*HANDSHAKE_UP) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HANDSHAKE_UP) GetImei() string {
	if m != nil && m.Imei != nil {
		return *m.Imei
	}
	return ""
}

func (m *HANDSHAKE_UP) GetImsi() string {
	if m != nil && m.Imsi != nil {
		return *m.Imsi
	}
	return ""
}

func (m *HANDSHAKE_UP) GetReport() string {
	if m != nil && m.Report != nil {
		return *m.Report
	}
	return ""
}

// 与服务器握手数据(下行)
type HANDSHAKE_DOWN struct {
	Port             *int32  `protobuf:"varint,1,req,name=port" json:"port,omitempty"`
	Ip               *string `protobuf:"bytes,2,req,name=ip" json:"ip,omitempty"`
	Time             *TIME   `protobuf:"bytes,3,req,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HANDSHAKE_DOWN) Reset()                    { *m = HANDSHAKE_DOWN{} }
func (m *HANDSHAKE_DOWN) String() string            { return proto.CompactTextString(m) }
func (*HANDSHAKE_DOWN) ProtoMessage()               {}
func (*HANDSHAKE_DOWN) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HANDSHAKE_DOWN) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *HANDSHAKE_DOWN) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *HANDSHAKE_DOWN) GetTime() *TIME {
	if m != nil {
		return m.Time
	}
	return nil
}

// 在线升级数据(上行)
type UPGRADE_UP struct {
	Version          *int32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	Index            *int32 `protobuf:"varint,2,req,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *UPGRADE_UP) Reset()                    { *m = UPGRADE_UP{} }
func (m *UPGRADE_UP) String() string            { return proto.CompactTextString(m) }
func (*UPGRADE_UP) ProtoMessage()               {}
func (*UPGRADE_UP) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UPGRADE_UP) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *UPGRADE_UP) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

// 在线升级数据(下行)
type UPGRADE_DOWN struct {
	Version          *int32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	Pieces           *int32 `protobuf:"varint,2,req,name=pieces" json:"pieces,omitempty"`
	Total            *int32 `protobuf:"varint,3,req,name=total" json:"total,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *UPGRADE_DOWN) Reset()                    { *m = UPGRADE_DOWN{} }
func (m *UPGRADE_DOWN) String() string            { return proto.CompactTextString(m) }
func (*UPGRADE_DOWN) ProtoMessage()               {}
func (*UPGRADE_DOWN) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UPGRADE_DOWN) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *UPGRADE_DOWN) GetPieces() int32 {
	if m != nil && m.Pieces != nil {
		return *m.Pieces
	}
	return 0
}

func (m *UPGRADE_DOWN) GetTotal() int32 {
	if m != nil && m.Total != nil {
		return *m.Total
	}
	return 0
}

// 开始充电指令
type START struct {
	Mins             *int32 `protobuf:"varint,1,req,name=mins" json:"mins,omitempty"`
	Time             *TIME  `protobuf:"bytes,2,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *START) Reset()                    { *m = START{} }
func (m *START) String() string            { return proto.CompactTextString(m) }
func (*START) ProtoMessage()               {}
func (*START) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *START) GetMins() int32 {
	if m != nil && m.Mins != nil {
		return *m.Mins
	}
	return 0
}

func (m *START) GetTime() *TIME {
	if m != nil {
		return m.Time
	}
	return nil
}

// 充电桩运行时基础数据
type BASE struct {
	Mode             *BASE_MODE   `protobuf:"varint,1,req,name=mode,enum=hooenergy.device.BASE_MODE" json:"mode,omitempty"`
	Rate             *int32       `protobuf:"varint,2,req,name=rate" json:"rate,omitempty"`
	Voltage          *int32       `protobuf:"varint,3,req,name=voltage" json:"voltage,omitempty"`
	Current          *int32       `protobuf:"varint,4,req,name=current" json:"current,omitempty"`
	Inputvol         *int32       `protobuf:"varint,5,req,name=inputvol" json:"inputvol,omitempty"`
	Light            *int32       `protobuf:"varint,6,opt,name=light" json:"light,omitempty"`
	Meter            *int32       `protobuf:"varint,7,opt,name=meter" json:"meter,omitempty"`
	Screen           *int32       `protobuf:"varint,8,opt,name=screen" json:"screen,omitempty"`
	Spd              *int32       `protobuf:"varint,9,opt,name=spd" json:"spd,omitempty"`
	Temp             *int32       `protobuf:"varint,10,opt,name=temp" json:"temp,omitempty"`
	Spear            *BASE_STATUS `protobuf:"varint,11,opt,name=spear,enum=hooenergy.device.BASE_STATUS" json:"spear,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *BASE) Reset()                    { *m = BASE{} }
func (m *BASE) String() string            { return proto.CompactTextString(m) }
func (*BASE) ProtoMessage()               {}
func (*BASE) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BASE) GetMode() BASE_MODE {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return BASE_STATUS_EXCEPTION
}

func (m *BASE) GetRate() int32 {
	if m != nil && m.Rate != nil {
		return *m.Rate
	}
	return 0
}

func (m *BASE) GetVoltage() int32 {
	if m != nil && m.Voltage != nil {
		return *m.Voltage
	}
	return 0
}

func (m *BASE) GetCurrent() int32 {
	if m != nil && m.Current != nil {
		return *m.Current
	}
	return 0
}

func (m *BASE) GetInputvol() int32 {
	if m != nil && m.Inputvol != nil {
		return *m.Inputvol
	}
	return 0
}

func (m *BASE) GetLight() int32 {
	if m != nil && m.Light != nil {
		return *m.Light
	}
	return 0
}

func (m *BASE) GetMeter() int32 {
	if m != nil && m.Meter != nil {
		return *m.Meter
	}
	return 0
}

func (m *BASE) GetScreen() int32 {
	if m != nil && m.Screen != nil {
		return *m.Screen
	}
	return 0
}

func (m *BASE) GetSpd() int32 {
	if m != nil && m.Spd != nil {
		return *m.Spd
	}
	return 0
}

func (m *BASE) GetTemp() int32 {
	if m != nil && m.Temp != nil {
		return *m.Temp
	}
	return 0
}

func (m *BASE) GetSpear() BASE_STATUS {
	if m != nil && m.Spear != nil {
		return *m.Spear
	}
	return BASE_STATUS_UNKNOWN
}

// 分时电价费率
type FEERATE struct {
	Fee              *float32 `protobuf:"fixed32,1,req,name=fee" json:"fee,omitempty"`
	Start            *uint32  `protobuf:"varint,2,req,name=start" json:"start,omitempty"`
	End              *uint32  `protobuf:"varint,3,req,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *FEERATE) Reset()                    { *m = FEERATE{} }
func (m *FEERATE) String() string            { return proto.CompactTextString(m) }
func (*FEERATE) ProtoMessage()               {}
func (*FEERATE) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FEERATE) GetFee() float32 {
	if m != nil && m.Fee != nil {
		return *m.Fee
	}
	return 0
}

func (m *FEERATE) GetStart() uint32 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

func (m *FEERATE) GetEnd() uint32 {
	if m != nil && m.End != nil {
		return *m.End
	}
	return 0
}

// 时段数据
type PERIOD struct {
	Start            *int32 `protobuf:"varint,1,req,name=start" json:"start,omitempty"`
	End              *int32 `protobuf:"varint,2,req,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PERIOD) Reset()                    { *m = PERIOD{} }
func (m *PERIOD) String() string            { return proto.CompactTextString(m) }
func (*PERIOD) ProtoMessage()               {}
func (*PERIOD) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PERIOD) GetStart() int32 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

func (m *PERIOD) GetEnd() int32 {
	if m != nil && m.End != nil {
		return *m.End
	}
	return 0
}

// 充电桩配置数据
type CONFIG struct {
	Category         *CONFIG_CATEGORY `protobuf:"varint,1,req,name=category,enum=hooenergy.device.CONFIG_CATEGORY" json:"category,omitempty"`
	Udid             *string          `protobuf:"bytes,2,req,name=udid" json:"udid,omitempty"`
	Maxvol           *int32           `protobuf:"varint,3,req,name=maxvol" json:"maxvol,omitempty"`
	Maxcur           *int32           `protobuf:"varint,4,req,name=maxcur" json:"maxcur,omitempty"`
	Minvol           *int32           `protobuf:"varint,5,req,name=minvol" json:"minvol,omitempty"`
	Ratevol          *int32           `protobuf:"varint,6,req,name=ratevol" json:"ratevol,omitempty"`
	Ratecur          *int32           `protobuf:"varint,7,req,name=ratecur" json:"ratecur,omitempty"`
	Version          *int32           `protobuf:"varint,8,opt,name=version" json:"version,omitempty"`
	Timefee          []*FEERATE       `protobuf:"bytes,9,rep,name=timefee" json:"timefee,omitempty"`
	Parkfee          *float32         `protobuf:"fixed32,10,opt,name=parkfee" json:"parkfee,omitempty"`
	Servicefee       *float32         `protobuf:"fixed32,11,opt,name=servicefee" json:"servicefee,omitempty"`
	Light            []*PERIOD        `protobuf:"bytes,12,rep,name=light" json:"light,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *CONFIG) Reset()                    { *m = CONFIG{} }
func (m *CONFIG) String() string            { return proto.CompactTextString(m) }
func (*CONFIG) ProtoMessage()               {}
func (*CONFIG) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CONFIG) GetCategory() CONFIG_CATEGORY {
	if m != nil && m.Category != nil {
		return *m.Category
	}
	return CONFIG_DEFAULT_UNKNOWN
}

func (m *CONFIG) GetUdid() string {
	if m != nil && m.Udid != nil {
		return *m.Udid
	}
	return ""
}

func (m *CONFIG) GetMaxvol() int32 {
	if m != nil && m.Maxvol != nil {
		return *m.Maxvol
	}
	return 0
}

func (m *CONFIG) GetMaxcur() int32 {
	if m != nil && m.Maxcur != nil {
		return *m.Maxcur
	}
	return 0
}

func (m *CONFIG) GetMinvol() int32 {
	if m != nil && m.Minvol != nil {
		return *m.Minvol
	}
	return 0
}

func (m *CONFIG) GetRatevol() int32 {
	if m != nil && m.Ratevol != nil {
		return *m.Ratevol
	}
	return 0
}

func (m *CONFIG) GetRatecur() int32 {
	if m != nil && m.Ratecur != nil {
		return *m.Ratecur
	}
	return 0
}

func (m *CONFIG) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *CONFIG) GetTimefee() []*FEERATE {
	if m != nil {
		return m.Timefee
	}
	return nil
}

func (m *CONFIG) GetParkfee() float32 {
	if m != nil && m.Parkfee != nil {
		return *m.Parkfee
	}
	return 0
}

func (m *CONFIG) GetServicefee() float32 {
	if m != nil && m.Servicefee != nil {
		return *m.Servicefee
	}
	return 0
}

func (m *CONFIG) GetLight() []*PERIOD {
	if m != nil {
		return m.Light
	}
	return nil
}

// 充电实时数据
type CHARGE struct {
	State            *CHARGE_STATE `protobuf:"varint,1,req,name=state,enum=hooenergy.device.CHARGE_STATE" json:"state,omitempty"`
	Duration         *int32        `protobuf:"varint,2,req,name=duration" json:"duration,omitempty"`
	Voltage          *int32        `protobuf:"varint,3,req,name=voltage" json:"voltage,omitempty"`
	Current          *int32        `protobuf:"varint,4,req,name=current" json:"current,omitempty"`
	Energy           *int32        `protobuf:"varint,5,req,name=energy" json:"energy,omitempty"`
	Amount           *float32      `protobuf:"fixed32,6,req,name=amount" json:"amount,omitempty"`
	Iccard           *string       `protobuf:"bytes,7,opt,name=iccard" json:"iccard,omitempty"`
	Remain           *float32      `protobuf:"fixed32,8,opt,name=remain" json:"remain,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *CHARGE) Reset()                    { *m = CHARGE{} }
func (m *CHARGE) String() string            { return proto.CompactTextString(m) }
func (*CHARGE) ProtoMessage()               {}
func (*CHARGE) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CHARGE) GetState() CHARGE_STATE {
	if m != nil && m.State != nil {
		return *m.State
	}
	return CHARGE_CHARGING_EXCEPTION
}

func (m *CHARGE) GetDuration() int32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *CHARGE) GetVoltage() int32 {
	if m != nil && m.Voltage != nil {
		return *m.Voltage
	}
	return 0
}

func (m *CHARGE) GetCurrent() int32 {
	if m != nil && m.Current != nil {
		return *m.Current
	}
	return 0
}

func (m *CHARGE) GetEnergy() int32 {
	if m != nil && m.Energy != nil {
		return *m.Energy
	}
	return 0
}

func (m *CHARGE) GetAmount() float32 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func (m *CHARGE) GetIccard() string {
	if m != nil && m.Iccard != nil {
		return *m.Iccard
	}
	return ""
}

func (m *CHARGE) GetRemain() float32 {
	if m != nil && m.Remain != nil {
		return *m.Remain
	}
	return 0
}

// 告警数据
type ALARM struct {
	Type             *ALARM_TYPE `protobuf:"varint,1,req,name=type,enum=hooenergy.device.ALARM_TYPE" json:"type,omitempty"`
	Start            *TIME       `protobuf:"bytes,2,req,name=start" json:"start,omitempty"`
	Block            *int32      `protobuf:"varint,3,req,name=block" json:"block,omitempty"`
	End              *TIME       `protobuf:"bytes,4,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ALARM) Reset()                    { *m = ALARM{} }
func (m *ALARM) String() string            { return proto.CompactTextString(m) }
func (*ALARM) ProtoMessage()               {}
func (*ALARM) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ALARM) GetType() ALARM_TYPE {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ALARM_NETWORK_FAULT
}

func (m *ALARM) GetStart() *TIME {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *ALARM) GetBlock() int32 {
	if m != nil && m.Block != nil {
		return *m.Block
	}
	return 0
}

func (m *ALARM) GetEnd() *TIME {
	if m != nil {
		return m.End
	}
	return nil
}

// 电表实时数据
type METER struct {
	Voltage          *int32 `protobuf:"varint,1,req,name=voltage" json:"voltage,omitempty"`
	Current          *int32 `protobuf:"varint,2,req,name=current" json:"current,omitempty"`
	Energy           *int32 `protobuf:"varint,3,req,name=energy" json:"energy,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *METER) Reset()                    { *m = METER{} }
func (m *METER) String() string            { return proto.CompactTextString(m) }
func (*METER) ProtoMessage()               {}
func (*METER) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *METER) GetVoltage() int32 {
	if m != nil && m.Voltage != nil {
		return *m.Voltage
	}
	return 0
}

func (m *METER) GetCurrent() int32 {
	if m != nil && m.Current != nil {
		return *m.Current
	}
	return 0
}

func (m *METER) GetEnergy() int32 {
	if m != nil && m.Energy != nil {
		return *m.Energy
	}
	return 0
}

// 交易记录数据
type TRADE struct {
	Tradeid          *string      `protobuf:"bytes,1,req,name=tradeid" json:"tradeid,omitempty"`
	Sequeue          *int32       `protobuf:"varint,2,req,name=sequeue" json:"sequeue,omitempty"`
	Type             *TRADE_TYPE  `protobuf:"varint,3,req,name=type,enum=hooenergy.device.TRADE_TYPE" json:"type,omitempty"`
	Iccard           *string      `protobuf:"bytes,4,req,name=iccard" json:"iccard,omitempty"`
	Remainpre        *float32     `protobuf:"fixed32,5,req,name=remainpre" json:"remainpre,omitempty"`
	Duration         *int32       `protobuf:"varint,6,req,name=duration" json:"duration,omitempty"`
	Amount           *float32     `protobuf:"fixed32,7,req,name=amount" json:"amount,omitempty"`
	Energy           *int32       `protobuf:"varint,8,req,name=energy" json:"energy,omitempty"`
	Energypre        *int32       `protobuf:"varint,9,req,name=energypre" json:"energypre,omitempty"`
	Energyend        *int32       `protobuf:"varint,10,req,name=energyend" json:"energyend,omitempty"`
	Paid             *int32       `protobuf:"varint,11,req,name=paid" json:"paid,omitempty"`
	Cause            *TRADE_CAUSE `protobuf:"varint,12,req,name=cause,enum=hooenergy.device.TRADE_CAUSE" json:"cause,omitempty"`
	Start            *TIME        `protobuf:"bytes,13,req,name=start" json:"start,omitempty"`
	End              *TIME        `protobuf:"bytes,14,req,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *TRADE) Reset()                    { *m = TRADE{} }
func (m *TRADE) String() string            { return proto.CompactTextString(m) }
func (*TRADE) ProtoMessage()               {}
func (*TRADE) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *TRADE) GetTradeid() string {
	if m != nil && m.Tradeid != nil {
		return *m.Tradeid
	}
	return ""
}

func (m *TRADE) GetSequeue() int32 {
	if m != nil && m.Sequeue != nil {
		return *m.Sequeue
	}
	return 0
}

func (m *TRADE) GetType() TRADE_TYPE {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return TRADE_USER_APP
}

func (m *TRADE) GetIccard() string {
	if m != nil && m.Iccard != nil {
		return *m.Iccard
	}
	return ""
}

func (m *TRADE) GetRemainpre() float32 {
	if m != nil && m.Remainpre != nil {
		return *m.Remainpre
	}
	return 0
}

func (m *TRADE) GetDuration() int32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *TRADE) GetAmount() float32 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func (m *TRADE) GetEnergy() int32 {
	if m != nil && m.Energy != nil {
		return *m.Energy
	}
	return 0
}

func (m *TRADE) GetEnergypre() int32 {
	if m != nil && m.Energypre != nil {
		return *m.Energypre
	}
	return 0
}

func (m *TRADE) GetEnergyend() int32 {
	if m != nil && m.Energyend != nil {
		return *m.Energyend
	}
	return 0
}

func (m *TRADE) GetPaid() int32 {
	if m != nil && m.Paid != nil {
		return *m.Paid
	}
	return 0
}

func (m *TRADE) GetCause() TRADE_CAUSE {
	if m != nil && m.Cause != nil {
		return *m.Cause
	}
	return TRADE_USER_ACTIVE
}

func (m *TRADE) GetStart() *TIME {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *TRADE) GetEnd() *TIME {
	if m != nil {
		return m.End
	}
	return nil
}

func init() {
	proto.RegisterType((*TIME)(nil), "hooenergy.device.TIME")
	proto.RegisterType((*HANDSHAKE_UP)(nil), "hooenergy.device.HANDSHAKE_UP")
	proto.RegisterType((*HANDSHAKE_DOWN)(nil), "hooenergy.device.HANDSHAKE_DOWN")
	proto.RegisterType((*UPGRADE_UP)(nil), "hooenergy.device.UPGRADE_UP")
	proto.RegisterType((*UPGRADE_DOWN)(nil), "hooenergy.device.UPGRADE_DOWN")
	proto.RegisterType((*START)(nil), "hooenergy.device.START")
	proto.RegisterType((*BASE)(nil), "hooenergy.device.BASE")
	proto.RegisterType((*FEERATE)(nil), "hooenergy.device.FEERATE")
	proto.RegisterType((*PERIOD)(nil), "hooenergy.device.PERIOD")
	proto.RegisterType((*CONFIG)(nil), "hooenergy.device.CONFIG")
	proto.RegisterType((*CHARGE)(nil), "hooenergy.device.CHARGE")
	proto.RegisterType((*ALARM)(nil), "hooenergy.device.ALARM")
	proto.RegisterType((*METER)(nil), "hooenergy.device.METER")
	proto.RegisterType((*TRADE)(nil), "hooenergy.device.TRADE")
	proto.RegisterEnum("hooenergy.device.BASE_MODE", BASE_MODE_name, BASE_MODE_value)
	proto.RegisterEnum("hooenergy.device.BASE_STATUS", BASE_STATUS_name, BASE_STATUS_value)
	proto.RegisterEnum("hooenergy.device.CONFIG_CATEGORY", CONFIG_CATEGORY_name, CONFIG_CATEGORY_value)
	proto.RegisterEnum("hooenergy.device.CHARGE_STATE", CHARGE_STATE_name, CHARGE_STATE_value)
	proto.RegisterEnum("hooenergy.device.ALARM_TYPE", ALARM_TYPE_name, ALARM_TYPE_value)
	proto.RegisterEnum("hooenergy.device.TRADE_TYPE", TRADE_TYPE_name, TRADE_TYPE_value)
	proto.RegisterEnum("hooenergy.device.TRADE_CAUSE", TRADE_CAUSE_name, TRADE_CAUSE_value)
}

func init() { proto.RegisterFile("hooenergy.device.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xdf, 0x6e, 0xe3, 0xc4,
	0x17, 0xfe, 0xd9, 0x71, 0x9c, 0xe4, 0xe4, 0xdf, 0xd4, 0xbb, 0xbf, 0x95, 0x11, 0x2c, 0xca, 0x9a,
	0x45, 0x94, 0xd5, 0x52, 0xa4, 0x72, 0x05, 0x48, 0x48, 0xae, 0x33, 0x4d, 0xac, 0x36, 0xb6, 0x65,
	0x3b, 0x5b, 0xf6, 0x2a, 0x32, 0xc9, 0xd0, 0x5a, 0xdb, 0xc4, 0xc1, 0x71, 0x56, 0x2d, 0x8f, 0x80,
	0xb8, 0xe2, 0x0d, 0x78, 0x05, 0xde, 0x86, 0xf7, 0xe0, 0x8e, 0x1b, 0x74, 0x66, 0xc6, 0x49, 0xda,
	0xd0, 0x8a, 0xcb, 0xf3, 0xf9, 0x9b, 0xf3, 0xf7, 0x3b, 0x33, 0x86, 0x67, 0x57, 0x59, 0xc6, 0x16,
	0x2c, 0xbf, 0xbc, 0x3d, 0x9a, 0xb1, 0xf7, 0xe9, 0x94, 0x1d, 0x2d, 0xf3, 0xac, 0xc8, 0x0c, 0x72,
	0x1f, 0xb7, 0xde, 0x81, 0x16, 0xbb, 0x23, 0x6a, 0xb4, 0x40, 0xbb, 0x65, 0x49, 0x6e, 0x2a, 0x3d,
	0xf5, 0xb0, 0x6a, 0xb4, 0xa1, 0x3a, 0xcf, 0x16, 0xc5, 0x95, 0xa9, 0x72, 0xb3, 0x09, 0x95, 0x59,
	0x72, 0x6b, 0x56, 0xb8, 0xd1, 0x02, 0xed, 0x2a, 0x5b, 0xe7, 0xa6, 0x56, 0x7e, 0x9a, 0xa7, 0x0b,
	0xb3, 0x5a, 0x1a, 0x2b, 0x36, 0x35, 0x75, 0x6e, 0x1c, 0x40, 0xa3, 0x48, 0xe7, 0x6c, 0x55, 0x24,
	0xf3, 0xa5, 0x59, 0xeb, 0x29, 0x87, 0x6d, 0xeb, 0x1b, 0x68, 0x0d, 0x6d, 0xaf, 0x1f, 0x0d, 0xed,
	0x33, 0x3a, 0x19, 0x07, 0xe8, 0x2a, 0x9d, 0xb3, 0x94, 0x07, 0x6d, 0x08, 0x6b, 0x95, 0xf2, 0x98,
	0x0d, 0xa3, 0x03, 0x7a, 0xce, 0x96, 0x59, 0x5e, 0x98, 0x95, 0x9e, 0x72, 0xd8, 0xb0, 0x02, 0xe8,
	0x6c, 0xcf, 0xf6, 0xfd, 0x0b, 0x0f, 0xf9, 0xfc, 0xbb, 0x48, 0x19, 0x40, 0x4d, 0x97, 0xf2, 0xec,
	0x4b, 0xd0, 0x30, 0x34, 0x4f, 0xb8, 0x79, 0xfc, 0xec, 0x68, 0xaf, 0x1b, 0x58, 0xb2, 0xf5, 0x1a,
	0x60, 0x1c, 0x0c, 0x42, 0xbb, 0xcf, 0x73, 0xe9, 0x42, 0xed, 0x3d, 0xcb, 0x57, 0x69, 0xb6, 0xd8,
	0xf6, 0x20, 0x5d, 0xcc, 0xd8, 0x8d, 0xe8, 0x81, 0xf5, 0x1d, 0xb4, 0x4a, 0x36, 0x8f, 0xbe, 0xc7,
	0xef, 0x80, 0xbe, 0x4c, 0xd9, 0x94, 0xad, 0x64, 0xd3, 0xda, 0x50, 0x2d, 0xb2, 0x22, 0xb9, 0x16,
	0x6d, 0xb3, 0xbe, 0x85, 0x6a, 0x14, 0xdb, 0x61, 0x8c, 0x69, 0xcf, 0xd3, 0xc5, 0x4a, 0x9e, 0x2a,
	0x53, 0x55, 0x7b, 0xca, 0x23, 0xa9, 0xfe, 0x5d, 0x01, 0xed, 0xc4, 0x8e, 0xa8, 0xf1, 0x39, 0x68,
	0xf3, 0x6c, 0xc6, 0xf8, 0xe1, 0xce, 0xf1, 0x87, 0xfb, 0x74, 0x64, 0x1d, 0x8d, 0xfc, 0x3e, 0x9f,
	0x68, 0x9e, 0x14, 0x4c, 0x66, 0x83, 0xe9, 0x66, 0xd7, 0x45, 0x72, 0xc9, 0xe4, 0x18, 0xbb, 0x50,
	0x9b, 0xae, 0xf3, 0x9c, 0x2d, 0x0a, 0x39, 0x49, 0x02, 0xf5, 0x74, 0xb1, 0x5c, 0x17, 0xef, 0xb3,
	0x6b, 0x39, 0xce, 0x36, 0x54, 0xaf, 0xd3, 0xcb, 0xab, 0xc2, 0xd4, 0x7b, 0x8a, 0x14, 0x05, 0x2b,
	0x58, 0xce, 0x87, 0xc9, 0xeb, 0x5d, 0x4d, 0x73, 0xc6, 0x16, 0x66, 0x9d, 0xdb, 0x38, 0xfc, 0xe5,
	0xcc, 0x6c, 0x70, 0xa3, 0x05, 0x5a, 0xc1, 0xe6, 0x4b, 0x13, 0xb8, 0xf5, 0x1a, 0xaa, 0xab, 0x25,
	0xaa, 0xab, 0xd9, 0x53, 0x0e, 0x3b, 0xc7, 0xcf, 0x1f, 0x48, 0x3b, 0x8a, 0xed, 0x78, 0x1c, 0x59,
	0xbf, 0x29, 0xa0, 0xf1, 0x0a, 0x9e, 0x02, 0x11, 0xd0, 0x84, 0x7e, 0xef, 0xd0, 0x20, 0x76, 0x7d,
	0x8f, 0xfc, 0xcf, 0xe8, 0x42, 0x53, 0xa2, 0x6e, 0xff, 0x9c, 0x12, 0xc5, 0x38, 0x80, 0xb6, 0x04,
	0x9c, 0xa1, 0x1d, 0x0e, 0x28, 0x51, 0x0d, 0x03, 0x3a, 0x12, 0x92, 0x33, 0x23, 0x95, 0x1d, 0x6f,
	0x8e, 0xef, 0x79, 0xd4, 0x89, 0x69, 0x9f, 0x68, 0xc6, 0x13, 0xe8, 0x6e, 0xd0, 0x51, 0x70, 0x4e,
	0x63, 0x4a, 0xaa, 0x3b, 0x1e, 0x4f, 0x5d, 0xcf, 0x8d, 0x86, 0x44, 0xb7, 0x7e, 0x06, 0x5d, 0x40,
	0xbb, 0xbe, 0xbd, 0x33, 0xcf, 0xbf, 0xc0, 0x9c, 0x4c, 0x78, 0x2a, 0xb1, 0xbe, 0x1b, 0xa1, 0xfb,
	0xb1, 0x77, 0xee, 0x3b, 0x67, 0x44, 0x31, 0xfe, 0x0f, 0x07, 0x3b, 0x51, 0x25, 0xac, 0xde, 0x09,
	0xeb, 0x09, 0xb0, 0x62, 0x3c, 0x03, 0xe3, 0xae, 0x17, 0x8e, 0x6b, 0xd6, 0x31, 0xd4, 0x4e, 0x29,
	0x0d, 0xed, 0x98, 0x62, 0x93, 0x7f, 0x64, 0x62, 0xfc, 0x2a, 0x0e, 0x64, 0x55, 0x24, 0x79, 0xc1,
	0x47, 0xdc, 0xc6, 0x6f, 0x6c, 0x31, 0xe3, 0xe3, 0x6d, 0x5b, 0x2f, 0x41, 0x0f, 0x68, 0xe8, 0xfa,
	0xfd, 0x2d, 0x4b, 0x29, 0x77, 0x14, 0x59, 0x42, 0xd4, 0x7f, 0xaa, 0xa0, 0x3b, 0xbe, 0x77, 0xea,
	0x0e, 0x8c, 0xaf, 0xa0, 0x3e, 0x4d, 0x0a, 0x76, 0x99, 0xe5, 0xb7, 0x52, 0x5d, 0x2f, 0xf6, 0xc7,
	0x24, 0xb8, 0x47, 0x8e, 0x1d, 0xd3, 0x81, 0x1f, 0xbe, 0xc5, 0x31, 0xaf, 0x67, 0xe9, 0x6c, 0xbb,
	0xb2, 0xf3, 0xe4, 0x06, 0xf5, 0x53, 0x29, 0x37, 0x62, 0x9e, 0xdc, 0x4c, 0x37, 0x77, 0x05, 0xda,
	0xe9, 0x62, 0xab, 0xaf, 0x2e, 0xd4, 0x50, 0xa1, 0x08, 0xe8, 0xbb, 0x00, 0x9e, 0xa8, 0x6d, 0x54,
	0x2b, 0x97, 0x4c, 0x88, 0xec, 0x15, 0xd4, 0x70, 0x5d, 0xb0, 0x07, 0x8d, 0x5e, 0xe5, 0xb0, 0x79,
	0xfc, 0xc1, 0x7e, 0x92, 0x65, 0xaf, 0xba, 0x50, 0x5b, 0x26, 0xf9, 0x3b, 0xe4, 0xa2, 0x0c, 0x51,
	0x15, 0xb0, 0x62, 0x39, 0x72, 0x10, 0x6b, 0x72, 0xec, 0xb3, 0x52, 0xe3, 0x2d, 0xee, 0xce, 0xdc,
	0x77, 0x27, 0xda, 0x68, 0x7d, 0x09, 0xf5, 0x4d, 0xd9, 0x4f, 0xa0, 0xdb, 0xa7, 0xa7, 0xf6, 0xf8,
	0x3c, 0xde, 0xd1, 0x40, 0x07, 0x60, 0x68, 0x7b, 0x03, 0xd7, 0x1b, 0x4c, 0x6c, 0x87, 0x28, 0xd6,
	0x5f, 0x0a, 0xe8, 0x42, 0x90, 0xc6, 0x17, 0x7c, 0x04, 0x45, 0xb9, 0xb6, 0x1f, 0xff, 0x4b, 0x63,
	0x39, 0x91, 0x6f, 0x00, 0xc5, 0x4d, 0x9c, 0xad, 0xf3, 0xa4, 0xc0, 0xb2, 0xff, 0xeb, 0xf6, 0x76,
	0x40, 0x17, 0x1e, 0x65, 0x6f, 0x3b, 0xa0, 0x27, 0xf3, 0x6c, 0xbd, 0x28, 0x78, 0x6b, 0x55, 0xb4,
	0xd3, 0xe9, 0x34, 0xc9, 0x67, 0x7c, 0x7b, 0xe5, 0xf5, 0x3a, 0x4f, 0x52, 0xd1, 0x58, 0xd5, 0x1a,
	0xf0, 0xeb, 0x29, 0xa6, 0x28, 0x42, 0x9e, 0x0c, 0xd6, 0x71, 0x6f, 0xed, 0x36, 0xb8, 0xef, 0x09,
	0x65, 0x6f, 0x80, 0xcd, 0xee, 0xa8, 0xd6, 0x1f, 0x15, 0xa8, 0xda, 0xe7, 0x76, 0x38, 0x32, 0x5e,
	0x81, 0x56, 0xdc, 0x2e, 0xcb, 0xa2, 0x3f, 0xda, 0x2f, 0x9a, 0xd3, 0x8e, 0xe2, 0xb7, 0x01, 0x35,
	0x3e, 0xdd, 0x95, 0xf2, 0x83, 0xf7, 0x20, 0x6a, 0xf9, 0x87, 0xeb, 0x6c, 0xfa, 0x4e, 0x76, 0xe1,
	0x13, 0xa1, 0x65, 0xed, 0xd1, 0xbb, 0xf3, 0x77, 0x15, 0x34, 0x1e, 0xe3, 0x00, 0xda, 0x1e, 0x8d,
	0x2f, 0xfc, 0xf0, 0x6c, 0xc2, 0x67, 0x47, 0x14, 0x83, 0x40, 0x2b, 0x72, 0x42, 0x4a, 0x3d, 0x89,
	0xa8, 0x58, 0xe6, 0x88, 0xc6, 0x34, 0x94, 0x40, 0x05, 0xc7, 0xea, 0x7a, 0xc1, 0x38, 0x9e, 0xf8,
	0x6f, 0x68, 0x48, 0x34, 0x24, 0x08, 0x7b, 0xec, 0xf5, 0x69, 0x48, 0x70, 0x14, 0x4d, 0x7f, 0x1c,
	0x6f, 0x18, 0x3a, 0x3a, 0x95, 0x80, 0xa0, 0xd4, 0x50, 0x2f, 0x3b, 0x94, 0x89, 0x33, 0x0e, 0x49,
	0xdd, 0x68, 0x43, 0x23, 0x0a, 0xfa, 0x32, 0x4e, 0x03, 0xdd, 0xc4, 0x74, 0x14, 0xf0, 0x9e, 0xd3,
	0x3e, 0x01, 0xec, 0x2f, 0x07, 0x22, 0xea, 0x45, 0x7e, 0x99, 0x4f, 0xd3, 0x68, 0x41, 0x9d, 0xc3,
	0xe7, 0xfe, 0x05, 0x69, 0x6d, 0x49, 0x01, 0xb5, 0xc3, 0xf2, 0x6c, 0x1b, 0x93, 0x76, 0x42, 0x3b,
	0x1a, 0x4e, 0xa2, 0xd8, 0x0f, 0x48, 0x87, 0xdb, 0x76, 0x58, 0x06, 0xeb, 0xa2, 0x8d, 0x77, 0x8b,
	0xb4, 0x89, 0xf5, 0x35, 0x54, 0x79, 0xd5, 0xbb, 0x42, 0x53, 0xee, 0x0b, 0x4d, 0xbd, 0x27, 0x34,
	0xf1, 0xae, 0xfd, 0xa2, 0x41, 0x35, 0xc6, 0x2b, 0x16, 0xa9, 0x45, 0x9e, 0xcc, 0x58, 0x3a, 0x93,
	0x0f, 0x7a, 0x17, 0x6a, 0x2b, 0xf6, 0xd3, 0x9a, 0xad, 0xcb, 0x47, 0xa8, 0x54, 0x44, 0xe5, 0x21,
	0x45, 0x70, 0x47, 0x42, 0x11, 0x5b, 0xc1, 0x6a, 0xdc, 0xd9, 0x01, 0x34, 0x84, 0x60, 0x97, 0x39,
	0xe3, 0x1a, 0x57, 0xef, 0xec, 0x89, 0x7e, 0x4f, 0xf5, 0xb5, 0x52, 0xf5, 0x32, 0xd9, 0x7a, 0xf9,
	0x4f, 0x22, 0x6c, 0x74, 0xd2, 0xb8, 0x0b, 0xa1, 0x92, 0xa0, 0xfc, 0xc3, 0x59, 0x26, 0xe9, 0xcc,
	0x6c, 0x72, 0xeb, 0x35, 0x54, 0xa7, 0xc9, 0x7a, 0xc5, 0xcc, 0x16, 0xcf, 0xfa, 0xf9, 0x43, 0x59,
	0x3b, 0xf6, 0x38, 0xda, 0x11, 0x72, 0xfb, 0x51, 0x21, 0x4b, 0xe5, 0x76, 0x1e, 0xfd, 0x41, 0x79,
	0x21, 0x85, 0xdb, 0x82, 0xfa, 0x38, 0xa2, 0xe1, 0xc4, 0x0e, 0x02, 0xa2, 0x18, 0x4d, 0xa8, 0x71,
	0xcb, 0x75, 0x88, 0x6a, 0xfd, 0xaa, 0x40, 0x55, 0x04, 0xee, 0x42, 0x53, 0x90, 0x9c, 0xd8, 0x7d,
	0x83, 0xcf, 0xe2, 0x13, 0xe8, 0x8a, 0x5b, 0x65, 0x67, 0x3b, 0x51, 0x2f, 0x21, 0x1d, 0xd9, 0xae,
	0x37, 0xf1, 0xfc, 0x78, 0x42, 0x3d, 0x7f, 0x3c, 0x18, 0x92, 0x0a, 0xbe, 0x69, 0x9b, 0x5d, 0x9f,
	0xc4, 0x34, 0x1c, 0x11, 0x0d, 0x65, 0x3c, 0x0e, 0x07, 0xd4, 0x73, 0xde, 0x0a, 0x84, 0x2b, 0x3d,
	0xb0, 0xe9, 0xd9, 0x24, 0xba, 0x70, 0x63, 0x67, 0x48, 0x74, 0x04, 0xfa, 0x76, 0x4c, 0x4b, 0xa0,
	0x76, 0xf2, 0x1c, 0xda, 0xd3, 0x6c, 0xbe, 0x2d, 0xe7, 0xa4, 0x15, 0xe0, 0x7f, 0xe7, 0x34, 0xbb,
	0x3e, 0x49, 0x56, 0xec, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xa1, 0x77, 0xea, 0x97, 0x0a,
	0x00, 0x00,
}
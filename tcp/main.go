package main

import (
	"encoding/base64"
	"fca/dal"
	"fca/libs"
	"fca/logic"
	"fmt"
	"net"
	"time"
)

var s string = "SwEAAABDAAAAAAABVFMyMDE2MTEyNDAxAAAAAAABCg80NjAyMDE2MDcyNTAwMDMSEDkxMzkyMDE2MDcyNTAwMDO7ZA=="
var Bukets map[string]*Bucket
var Logic logic.LogicBase

func main() {
	libs.NewLogger()
	Logic = logic.LogicBase{}
	dal.InitDb()
	logic.InitLogic()
	go InitHttp()
	/*
		b, _ := base64.StdEncoding.DecodeString(s)
		device := &logic.DeviceProto{}
		device.ParseHead(b)
		pb, err := logic.Handshake_down(device.Content)
		pb2, err := logic.Handshake_up(device.Content)
		fmt.Printf("device:%v\n", device)
		fmt.Printf("pb:%v\n", pb)
		fmt.Printf("pb2:%v\n", &pb2)
		fmt.Printf("err:%v\n", err)
		 libs.NewLogger()
		b, _ := base64.StdEncoding.DecodeString(s)
		crc := ^libs.Crc16(b[0 : len(b)-2])
		fmt.Printf("crc:%v\n", crc)
		device := &logic.DeviceProto{}
		device.ParseHead(b)
		fmt.Printf("device:%v\n", device)
		//fmt.Printf("%v\n", b)
		Inittcp()
		device.Parse()
	*/
	Bukets = make(map[string]*Bucket)
	addres := make([]string, 0)
	addres = append(addres, "0.0.0.0:8998")
	InitTCP(addres, 10)
	c := make(chan int)
	<-c
}

const (
	ip   = ""
	port = 8998
)

func Inittcp() {
	addres := make([]string, 0)
	addres = append(addres, "0.0.0.0:8998")
	InitTCP(addres, 10)
	/*
		listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), port, ""})
		if err != nil {
			fmt.Println("监听端口失败:", err.Error())
			return
		}
		fmt.Println("已初始化连接，等待客户端连接...")
		go Server(listen)
	*/
}

func Server(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			panic(err)
			fmt.Println("接受客户端连接异常:", err.Error())
			continue
		}
		fmt.Println("客户端连接来自:", conn.RemoteAddr().String())
		down := logic.DeviceProto{}
		bytes, _ := down.Handshake_down(8998, "119.23.22.219", time.Now())

		go func() {
			device := logic.DeviceProto{}
			data := make([]byte, 128)
			for {
				i, err := conn.Read(data)
				fmt.Printf("device:%d，%v\n", i, data)
				fmt.Println(base64.StdEncoding.EncodeToString(data[0:i]))
				b := []byte(base64.StdEncoding.EncodeToString(data[0:i]))
				fmt.Printf("%v\n", b)
				device.ParseHead(data[0:i])
				bb := down.BuildHead(device.UDType, device.UDID, 1, bytes)
				conn.Write(bb)
				fmt.Printf("device:%v\n", device)
				fmt.Println("客户端发来数据:", string(data[0:i]))
				if err != nil {
					fmt.Println("读取客户端数据错误:", err.Error())
					break
				}
				conn.Write([]byte{'f', 'i', 'n', 'i', 's', 'h'})
			}

		}()

	}
}

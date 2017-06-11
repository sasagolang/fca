package main

import (
	"fca/libs"
	"fca/libs/bufio"
	"strconv"

	"fca/libs/bytes"
	"fmt"
	"net"
	"sync"
)

var cbLock sync.RWMutex

func InitTCP(addrs []string, accept int) (err error) {
	var (
		bind     string
		listener *net.TCPListener
		addr     *net.TCPAddr
	)
	for _, bind = range addrs {
		if addr, err = net.ResolveTCPAddr("tcp4", bind); err != nil {

			return
		}
		if listener, err = net.ListenTCP("tcp4", addr); err != nil {

			return
		}

		// split N core accept
		for i := 0; i < accept; i++ {
			go acceptTCP(listener)
		}
	}
	return
}

func acceptTCP(lis *net.TCPListener) {

	var (
		conn *net.TCPConn
		err  error
		r    int
	)
	for {
		if conn, err = lis.AcceptTCP(); err != nil {
			// if listener close then return
			fmt.Println(err.Error())
			return
		}
		if err = conn.SetKeepAlive(true); err != nil {
			fmt.Println(err.Error())
			return
		}
		if err = conn.SetReadBuffer(1024); err != nil {
			fmt.Println(err.Error())
			return
		}
		if err = conn.SetWriteBuffer(1024); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(conn.RemoteAddr().String())
		go serveTCP(conn, r)
		if r++; r == 100000 {
			r = 0
		}
	}
}

func serveTCP(conn *net.TCPConn, r int) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常

		if err := recover(); err != nil {

			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
			return
		}

	}()

	var ch = NewChannel("")
	ch.Reader = bufio.Reader{}
	ch.Writer = bufio.Writer{}
	var (
		rr = &ch.Reader
		wr = &ch.Writer
		p  = &libs.Proto{}
	)
	rb := bytes.NewBuffer(10000000)
	wb := bytes.NewBuffer(10000000)

	ch.Reader.ResetBuffer(conn, rb.Bytes())
	ch.Writer.ResetBuffer(conn, wb.Bytes())
	re := p.ReadTCP(rr)
	//添加到buket
	//fmt.Printf("serveTCP:%v\n", re)
	ch.UUID = re.UDID
	//ch.UUID = strings.Replace(ch.UUID, string([]byte{0}), "", -1)
	//ch.UUID = strings.Replace(ch.UUID, "\n", "", -1)
	ch.UDType = re.UDType
	fmt.Printf("链接:%v\n", ch.UUID)
	pole := Logic.GetPole(ch.UUID)
	ch.EPName = strconv.Itoa(int(pole.ElectricPile.ID))
	/*cbLock.Lock()
	if bucket,ok:=Bukets[strconv.Itoa( int(pole.ElectricPile.ID))];ok{
		bucket.Put(ch.UUID,ch)
	}else {
		Bukets[strconv.Itoa( int(pole.ElectricPile.ID))],_=NewBucket(strconv.Itoa( int(pole.ElectricPile.ID)))
		bucket=Bukets[strconv.Itoa( int(pole.ElectricPile.ID))]
		bucket.Put(ch.UUID,ch)
	}
	cbLock.Unlock()*/
	InBucket(ch)
	go dispatchTCP(ch, rr, wr)
	for {
		//var p1 = &libs.Proto{}
		p1 := p.ReadTCP(rr)
		Logic.ReceivedMsg(p1, ch.signal)
		//	fmt.Printf("serveTCP:%v\n", p1)
	}
}
func dispatchTCP(ch *Channel, rr *bufio.Reader, wr *bufio.Writer) {
	for {

		p := <-ch.signal
		p.UDID = ch.UUID
		p.UDType = ch.UDType
		//	fmt.Printf("dispatchTCP:%v\n", p)
		p.WriteTo(wr)
		wr.Flush()
	}

}

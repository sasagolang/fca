package main

import (
	"fmt"
	"strconv"
	"sync"
)

var cBuketsLock sync.RWMutex

func InBucket(ch *Channel) {
	//	fmt.Printf("Bukets_1:%v\n", Bukets)
	var (
		b  *Bucket
		ok bool
	)
	cBuketsLock.Lock()
	defer cBuketsLock.Unlock()
	fmt.Printf("InBucket.epname:%v,%v\n", ch.EPName, Bukets[ch.EPName])
	if b, ok = Bukets[ch.EPName]; !ok {
		fmt.Printf("InBucket:%v,%v\n", ch.EPName, Bukets[ch.EPName])
		b, _ = NewBucket(ch.EPName)

		Bukets[ch.EPName] = b

	}
	b.Put(ch.UUID, ch)
	//	fmt.Printf("Bukets_2:%v\n", Bukets)
}
func GetChannelByUUID(uuid string) *Channel {

	pole := Logic.GetPole(uuid)

	epname := strconv.Itoa(int(pole.ElectricPile.ID))
	fmt.Printf("epname:%v,%v\n", epname, Bukets[epname])
	if b, ok := Bukets[epname]; ok {
		fmt.Printf("GetChannelByUUID:%v\n", b)
		return b.Channel(uuid)
	}
	return nil
}

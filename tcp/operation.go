package main

import "sync"

var  cBuketsLock    sync.RWMutex
func InBucket(ch *Channel) {
	var (
		b  *Bucket
		ok bool
	)
	if b, ok = Bukets[ch.EPName]; !ok {
		b, _ = NewBucket(ch.EPName)
		cBuketsLock.Lock()
		Bukets[ch.EPName]=b
		cBuketsLock.Unlock()
	}
	b.Put(ch.UUID,ch)
}
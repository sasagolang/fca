package main

import "sync"
import "fmt"

type Bucket struct {
	cLock sync.RWMutex
	name  string
	chs   map[string]*Channel
}

func NewBucket(name string) (b *Bucket, err error) {
	b = &Bucket{}
	b.name = name
	b.chs = make(map[string]*Channel)
	return
}
func (b *Bucket) Put(key string, ch *Channel) (err error) {
	b.cLock.Lock()
	defer b.cLock.Unlock()
	/*if _, ok := b.chs[key]; !ok {
		b.chs[key] = ch
	}*/
	b.chs[key] = ch
	return
}
func (b *Bucket) Del(key string) (err error) {
	b.cLock.Lock()
	defer b.cLock.Unlock()
	if _, ok := b.chs[key]; ok {
		delete(b.chs, key)
	}
	return
}
func (b *Bucket) Channel(key string) (ch *Channel) {
	fmt.Printf("Channel:%v\n", &b)
	ch = b.chs[key]
	return
}

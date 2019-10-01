package sample

import (
	"sync"
	"sync/atomic"
)

var index int64

var aindex []*int64

var mutex sync.Mutex

func asmStackIndex() (stack uint32)

func init() {
	aindex = make([]*int64, 256)
	slice := aindex
	for i := range slice {
		slice[i] = new(int64)
	}
}

func Version0(n int64) {
	// Just empty function
}

func Version1(n int64) {
	index += n
}

func Version2(n int64) {
	mutex.Lock()
	index += n
	mutex.Unlock()
}

func Version3(n int64) {
	atomic.AddInt64(&index, n)
}

func Version4(i int, n int64) {
	atomic.AddInt64(aindex[asmStackIndex()], n)
}

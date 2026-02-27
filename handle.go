//go:build !cgo

package nocgo

import (
	"sync"
	"sync/atomic"
)

var handleKey atomic.Uintptr
var handles sync.Map

type Handle uintptr

func NewHandle(value any) Handle {
	key := handleKey.Add(1)
	if key == 0 {
		panic("nocgo: ran out of handle space")
	}
	handles.Store(key, value)
	return Handle(key)
}

func (h Handle) Value() any {
	value, ok := handles.Load(uintptr(h))
	if !ok {
		panic("nocgo: misuse of an invalid handle")
	}
	return value
}

func (h Handle) Delete() {
	_, ok := handles.LoadAndDelete(uintptr(h))
	if !ok {
		panic("nocgo: misuse of an invalid handle")
	}
}

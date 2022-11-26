package reflect_utils

import (
	"fmt"
	"reflect"
	"unsafe"
)

// SafeClose 安全的关闭channel，如果已经被关闭的话会recover住panic
// 不推荐使用此方法，一个好的结构清晰的代码不应该发生channel被重复关闭的情况，它只会在发送侧关闭一次才对
func SafeClose[T any](channel chan T) bool {

	if channel == nil {
		return true
	}

	defer func() {
		_ = recover()
	}()

	close(channel)
	return true
}

// IsClosed 判断channel是否已经被关闭
func IsClosed[T any](channel chan T) bool {

	// 传入的参数非法的话就认为是已经关闭的channel
	if channel == nil || reflect.TypeOf(channel).Kind() != reflect.Chan {
		return true
	}

	// get interface value pointer, from cgo_export
	// typedef struct { void *t; void *v; } GoInterface;
	// then get channel real pointer
	cptr := *(*uintptr)(unsafe.Pointer(
		unsafe.Pointer(uintptr(unsafe.Pointer(&channel)) + unsafe.Sizeof(uint(0))),
	))

	// this function will return true if chan.closed > 0
	// see hchan on https://github.com/golang/go/blob/master/src/runtime/chan.go
	// type hchan struct {
	// qcount   uint           // total data in the queue
	// dataqsiz uint           // size of the circular queue
	// buf      unsafe.Pointer // points to an array of dataqsiz elements
	// elemsize uint16
	// closed   uint32
	// **

	// qcount + dataqsiz
	cptr += unsafe.Sizeof(uint(0)) * 2
	// buf
	cptr += unsafe.Sizeof(uintptr(0))
	// elemsize
	cptr += unsafe.Sizeof(uint16(0))
	t := unsafe.Pointer(cptr)
	fmt.Println(t)
	return *(*uint32)(t) > 0
}



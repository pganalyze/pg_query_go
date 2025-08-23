//go:build !go1.20
// +build !go1.20

package parser

import (
	"reflect"
	"unsafe"
)

func toBytes(s string) (b []byte) {
	pb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	ps := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pb.Data = ps.Data
	pb.Len = ps.Len
	pb.Cap = ps.Len

	return b
}

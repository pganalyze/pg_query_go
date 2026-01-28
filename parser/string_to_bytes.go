//go:build go1.20
// +build go1.20

package parser

import "unsafe"

func toBytes(s string) (b []byte) {
	if s == "" {
		return nil
	}
	
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

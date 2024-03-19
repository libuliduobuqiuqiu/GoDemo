package godemo

import (
	"bytes"
	"unsafe"
)

var strColon = []byte(":")
var strSlash = []byte("/")
var strStar = []byte("*")

func CountParams(path string) uint16 {
	s := unsafe.Slice(unsafe.StringData(path), len(path))

	n := uint16(bytes.Count(s, strColon))
	n += uint16(bytes.Count(s, strSlash))
	return n
}

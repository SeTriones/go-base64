package base64

/*
#include "base64.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

var (
	ErrBadFormat = errors.New("bad base64 string format")
)

func Encode(input []byte) string {
	l := len(input)
	var output *C.char
	var outputLen C.int
	C.base64_encode((*C.uint8_t)(unsafe.Pointer(&input[0])), C.int(l), &output, &outputLen)
	defer C.free(unsafe.Pointer(output))
	return C.GoStringN(output, outputLen)
}

func Decode(input string) ([]byte, error) {
	return nil, nil
}

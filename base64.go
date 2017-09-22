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
	strSlice := []byte(input)
	inputLen := len(input)
	if inputLen%4 != 0 {
		return nil, ErrBadFormat
	}
	outputLen := inputLen / 4 * 3
	for idx := 1; idx <= 2; idx++ {
		if input[inputLen-idx] == 0x3d {
			outputLen = outputLen - 1
		}
	}
	outbuf := make([]byte, outputLen, outputLen)
	cret := C.base64_decode((*C.char)(unsafe.Pointer(&strSlice[0])), C.int(inputLen), (*C.uint8_t)(unsafe.Pointer(&outbuf[0])), C.int(outputLen))
	ret := int(cret)
	if ret < 0 {
		return nil, ErrBadFormat
	} else {
		return outbuf, nil
	}

}

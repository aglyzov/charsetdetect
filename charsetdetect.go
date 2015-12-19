// Package charsetdetect provides minimal cgo bindings for libcharsetdetect (included)
//
// Source code and project home:
// https://github.com/aglyzov/go-charsetdetect
//
package charsetdetect

/*
#cgo CFLAGS: -Ilibcharsetdetect
#cgo LDFLAGS: -Llibcharsetdetect/build -lcharsetdetect

#include <stdlib.h>
#include <charsetdetect.h>

const csd_t CSD_FAIL = (csd_t)(-1);
*/
import "C"

import (
	"unsafe"
	"errors"
)


func DetectCharset(text []byte) (charset string, err error) {
	var csd = C.csd_open()

	if csd == C.CSD_FAIL {
		err = errors.New("libcharsetdetect: init failed")
		return
	}

	C.csd_consider(csd, (*C.char)(unsafe.Pointer(&text[0])), C.int(len(text)))

	c_str := C.csd_close(csd)
	if c_str != (*C.char)(nil) {
		charset = C.GoString(c_str)
	}

	return
}


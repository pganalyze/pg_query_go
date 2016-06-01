package parser

/*
#cgo CFLAGS: -Iinclude
#cgo LDFLAGS: -fstack-protector
#include "pg_query.h"
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"sync"
	"unsafe"
)

func init() {
	C.pg_query_init()
}

var parseMutex sync.Mutex

// ParseToJSON - Parses the given SQL statement into an AST (JSON format)
func ParseToJSON(input string) (result string, err error) {
	inputC := C.CString(input)
	defer C.free(unsafe.Pointer(inputC))

	// Due to Postgres' forked model we need to prevent concurrent runs of the native code
	parseMutex.Lock()
	resultC := C.pg_query_parse(inputC)
	parseMutex.Unlock()

	defer C.pg_query_free_parse_result(resultC)

	if resultC.error != nil {
		errMessage := C.GoString(resultC.error.message)
		err = errors.New(errMessage)
		return
	}

	result = C.GoString(resultC.parse_tree)

	return
}

// Normalize the passed SQL statement to replace constant values with ? characters
func Normalize(input string) (result string, err error) {
	inputC := C.CString(input)
	defer C.free(unsafe.Pointer(inputC))

	resultC := C.pg_query_normalize(inputC)
	defer C.pg_query_free_normalize_result(resultC)

	if resultC.error != nil {
		errMessage := C.GoString(resultC.error.message)
		err = errors.New(errMessage)
		return
	}

	result = C.GoString(resultC.normalized_query)

	return
}

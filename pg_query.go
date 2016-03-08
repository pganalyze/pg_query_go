package pg_query

// Note(LukasFittl): This needs Go 1.5 for $SRCDIR support, see
// https://github.com/golang/go/commit/131758183f7dc2610af489da3a7fcc4d30c6bc48

/*
#cgo CFLAGS: -I${SRCDIR}/tmp/libpg_query
#cgo LDFLAGS: -L${SRCDIR}/tmp/libpg_query -lpg_query -fstack-protector
#include <pg_query.h>
#include <stdlib.h>
*/
import "C"

import (
	"encoding/json"
	"errors"
	"runtime/debug"
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

// Parse the given SQL statement into an AST (native Go structs)
func Parse(input string) (tree ParsetreeList, err error) {
	jsonTree, err := ParseToJSON(input)
	if err != nil {
		return
	}

	// JSON unmarshalling can panic in edge cases we don't support yet. This is
	// still a *bug that needs to be fixed*, but this way the caller can expect an
	// error to be returned always, instead of a panic
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			err = r.(error)
		}
	}()

	err = json.Unmarshal([]byte(jsonTree), &tree)
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

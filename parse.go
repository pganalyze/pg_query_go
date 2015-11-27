package pg_query

// Note(LukasFittl): This needs Go 1.5 for $SRCDIR support, see
// https://github.com/golang/go/commit/131758183f7dc2610af489da3a7fcc4d30c6bc48

/*
#cgo CFLAGS: -I${SRCDIR}/tmp/libpg_query-master
#cgo LDFLAGS: -L${SRCDIR}/tmp/libpg_query-master -lpg_query -fstack-protector
#include <pg_query.h>
#include <stdlib.h>
*/
import "C"

import (
	"encoding/json"
	"unsafe"
)

func init() {
	C.pg_query_init()
}

// ParseToJSON - Parses the given SQL statement into an AST (JSON format)
func ParseToJSON(input string) string {
	inputC := C.CString(input)
	defer C.free(unsafe.Pointer(inputC))

	resultC := C.pg_query_parse(inputC)
	defer C.pg_query_free_parse_result(resultC)

	result := C.GoString(resultC.parse_tree)

	return result
}

// Parse - Parses the given SQL statement into an AST (native Go structs)
func Parse(input string) (tree ParsetreeList, err error) {
	jsonTree := ParseToJSON(input)
	err = json.Unmarshal([]byte(jsonTree), &tree)
	return
}

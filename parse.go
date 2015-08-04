package pg_query_go

// Note(LukasFittl): This needs Go 1.5 for $SRCDIR support, see
// https://github.com/golang/go/commit/131758183f7dc2610af489da3a7fcc4d30c6bc48

/*
#cgo CFLAGS: -I${SRCDIR} -I ${SRCDIR}/postgres/src/include
#cgo LDFLAGS: -L${SRCDIR} -lpg_query -fstack-protector
#cgo LDFLAGS: -Wl,-undefined,dynamic_lookup -Wl,-multiply_defined,suppress
#include <pg_query.h>
*/
import "C"

import "unsafe"

func Parse(input string) string {
  C.pg_query_init()

  input_c := C.CString(input)
  defer C.free(unsafe.Pointer(input_c))

  result_c := C.pg_query_raw_parse(input_c)

  result := C.GoString(result_c)

  return result
}

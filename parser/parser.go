package parser

/*
#cgo CFLAGS: -Iinclude -g -fstack-protector
#cgo LDFLAGS:
#include "pg_query.h"
#include <stdlib.h>

// Avoid complexities dealing with C structs in Go
PgQueryDeparseResult pg_query_deparse_protobuf_direct_args(void* data, unsigned int len) {
	PgQueryProtobuf p;
	p.data = (char *) data;
	p.len = len;
	return pg_query_deparse_protobuf(p);
}
*/
import "C"

import (
	"errors"
	"unsafe"
)

func init() {
	C.pg_query_init()
}

// ParseToJSON - Parses the given SQL statement into a parse tree (JSON format)
func ParseToJSON(input string) (result string, err error) {
	inputC := C.CString(input)
	defer C.free(unsafe.Pointer(inputC))

	resultC := C.pg_query_parse(inputC)

	defer C.pg_query_free_parse_result(resultC)

	if resultC.error != nil {
		errMessage := C.GoString(resultC.error.message)
		err = errors.New(errMessage)
		return
	}

	result = C.GoString(resultC.parse_tree)

	return
}

// ParseToProtobuf - Parses the given SQL statement into a parse tree (Protobuf format)
func ParseToProtobuf(input string) (result []byte, err error) {
	inputC := C.CString(input)
	defer C.free(unsafe.Pointer(inputC))

	resultC := C.pg_query_parse_protobuf(inputC)

	defer C.pg_query_free_protobuf_parse_result(resultC)

	if resultC.error != nil {
		errMessage := C.GoString(resultC.error.message)
		err = errors.New(errMessage)
		return
	}

	result = []byte(C.GoStringN(resultC.parse_tree.data, C.int(resultC.parse_tree.len)))

	return
}

// DeparseFromProtobuf - Deparses the given Protobuf format parse tree into a SQL statement
func DeparseFromProtobuf(input []byte) (result string, err error) {
	inputC := C.CBytes(input)
	defer C.free(inputC)

	resultC := C.pg_query_deparse_protobuf_direct_args(inputC, C.uint(len(input)))

	defer C.pg_query_free_deparse_result(resultC)

	if resultC.error != nil {
		errMessage := C.GoString(resultC.error.message)
		err = errors.New(errMessage)
		return
	}

	result = C.GoString(resultC.query)

	return
}

// ParsePlPgSqlToJSON - Parses the given PL/pgSQL function statement into a parse tree (JSON format)
func ParsePlPgSqlToJSON(input string) (result string, err error) {
	inputC := C.CString(input)
	defer C.free(unsafe.Pointer(inputC))

	resultC := C.pg_query_parse_plpgsql(inputC)

	defer C.pg_query_free_plpgsql_parse_result(resultC)

	if resultC.error != nil {
		errMessage := C.GoString(resultC.error.message)
		err = errors.New(errMessage)
		return
	}

	result = C.GoString(resultC.plpgsql_funcs)

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

// Fingerprint - Fingerprint the passed SQL statement using the C extension
func Fingerprint(input string) (result string, err error) {
	inputC := C.CString(input)
	defer C.free(unsafe.Pointer(inputC))

	resultC := C.pg_query_fingerprint(inputC)
	defer C.pg_query_free_fingerprint_result(resultC)

	if resultC.error != nil {
		errMessage := C.GoString(resultC.error.message)
		err = errors.New(errMessage)
		return
	}

	result = C.GoString(resultC.fingerprint_str)

	return
}

package parser

/*
#cgo CFLAGS: -Iinclude -g -fstack-protector -std=gnu99
#cgo LDFLAGS:
#include "pg_query.h"
#include "xxhash.h"
#include <stdlib.h>

// Avoid complexities dealing with C structs in Go
PgQueryDeparseResult pg_query_deparse_protobuf_direct_args(void* data, unsigned int len) {
	PgQueryProtobuf p;
	p.data = (char *) data;
	p.len = len;
	return pg_query_deparse_protobuf(p);
}

// Avoid inconsistent type behaviour in xxhash library
uint64_t pg_query_hash_xxh3_64(void *data, size_t len, size_t seed) {
	return XXH3_64bits_withSeed(data, len, seed);
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

// Scans the given SQL statement into a protobuf ScanResult
func ScanToProtobuf(input string) (result []byte, err error) {
	inputC := C.CString(input)
	defer C.free(unsafe.Pointer(inputC))

	resultC := C.pg_query_scan(inputC)
	defer C.pg_query_free_scan_result(resultC)

	if resultC.error != nil {
		errMessage := C.GoString(resultC.error.message)
		err = errors.New(errMessage)
		return
	}

	result = []byte(C.GoStringN(resultC.pbuf.data, C.int(resultC.pbuf.len)))

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

// FingerprintToUInt64 - Fingerprint the passed SQL statement using the C extension and returns result as uint64
func FingerprintToUInt64(input string) (result uint64, err error) {
	inputC := C.CString(input)
	defer C.free(unsafe.Pointer(inputC))

	resultC := C.pg_query_fingerprint(inputC)
	defer C.pg_query_free_fingerprint_result(resultC)

	if resultC.error != nil {
		errMessage := C.GoString(resultC.error.message)
		err = errors.New(errMessage)
		return
	}

	// https://github.com/golang/go/issues/29878
	result = *(*uint64)(unsafe.Pointer(&resultC.fingerprint))

	return
}

// FingerprintToHexStr - Fingerprint the passed SQL statement using the C extension and returns result as hex string
func FingerprintToHexStr(input string) (result string, err error) {
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

// HashXXH3_64 - Helper method to run XXH3 hash function (64-bit variant) on the given bytes, with the specified seed
func HashXXH3_64(input []byte, seed uint64) (result uint64) {
	inputC := C.CBytes(input)
	defer C.free(inputC)

	res := C.pg_query_hash_xxh3_64(inputC, C.size_t(len(input)), C.size_t(seed))

	// https://github.com/golang/go/issues/29878
	result = *(*uint64)(unsafe.Pointer(&res))

	return
}

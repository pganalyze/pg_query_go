package pg_query

import (
	"io"
	"strconv"
)

func (node Value) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Type)))
	io.WriteString(ctx.hash, node.Str)
	io.WriteString(ctx.hash, strconv.Itoa(node.Ival))
}

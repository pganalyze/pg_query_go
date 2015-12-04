// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node SetOp) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SETOP")
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Cmd)))
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Strategy)))
}

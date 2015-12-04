// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node Hash) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "HASH")
	io.WriteString(ctx.hash, strconv.FormatBool(node.SkewInherit))
}

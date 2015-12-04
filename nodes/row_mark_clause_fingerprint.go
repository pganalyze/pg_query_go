// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node RowMarkClause) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ROWMARKCLAUSE")
	io.WriteString(ctx.hash, strconv.FormatBool(node.NoWait))
	io.WriteString(ctx.hash, strconv.FormatBool(node.PushedDown))
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Strength)))
}

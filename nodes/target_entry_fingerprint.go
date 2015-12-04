// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node TargetEntry) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "TARGETENTRY")

	if node.Expr != nil {
		node.Expr.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Resjunk))

	if node.Resname != nil {
		io.WriteString(ctx.hash, *node.Resname)
	}
}

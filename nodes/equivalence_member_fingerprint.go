// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node EquivalenceMember) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "EQUIVALENCEMEMBER")

	if node.EmExpr != nil {
		node.EmExpr.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.EmIsChild))
	io.WriteString(ctx.hash, strconv.FormatBool(node.EmIsConst))
}

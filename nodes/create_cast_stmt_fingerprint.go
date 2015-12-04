// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node CreateCastStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATECASTSTMT")
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Context)))

	if node.Func != nil {
		node.Func.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Inout))

	if node.Sourcetype != nil {
		node.Sourcetype.Fingerprint(ctx)
	}

	if node.Targettype != nil {
		node.Targettype.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateCastStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateCastStmt")
	if node.Func != nil {
		node.Func.Fingerprint(ctx)
	}

	if node.Sourcetype != nil {
		node.Sourcetype.Fingerprint(ctx)
	}

	if node.Targettype != nil {
		node.Targettype.Fingerprint(ctx)
	}
}

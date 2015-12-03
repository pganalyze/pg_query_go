// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ReindexStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ReindexStmt")
	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}
}

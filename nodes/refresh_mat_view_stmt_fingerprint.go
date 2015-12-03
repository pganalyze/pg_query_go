// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node RefreshMatViewStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RefreshMatViewStmt")
	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}
}

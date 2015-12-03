// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ClusterStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ClusterStmt")
	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node TableLikeClause) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "TABLELIKECLAUSE")

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}
}

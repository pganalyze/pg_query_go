// Auto-generated - DO NOT EDIT

package pg_query

func (node TableLikeClause) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("TABLELIKECLAUSE")

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}
}

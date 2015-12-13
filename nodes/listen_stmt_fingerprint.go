// Auto-generated - DO NOT EDIT

package pg_query

func (node ListenStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ListenStmt")

	if node.Conditionname != nil {
		ctx.WriteString(*node.Conditionname)
	}
}

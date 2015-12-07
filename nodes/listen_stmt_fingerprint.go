// Auto-generated - DO NOT EDIT

package pg_query

func (node ListenStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("LISTENSTMT")

	if node.Conditionname != nil {
		ctx.WriteString(*node.Conditionname)
	}
}

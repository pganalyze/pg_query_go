// Auto-generated - DO NOT EDIT

package pg_query

func (node UnlistenStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("UnlistenStmt")

	if node.Conditionname != nil {
		ctx.WriteString(*node.Conditionname)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

func (node VariableShowStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("SHOW")

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}
}

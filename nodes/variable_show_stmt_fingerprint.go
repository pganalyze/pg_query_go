// Auto-generated - DO NOT EDIT

package pg_query

func (node VariableShowStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("VariableShowStmt")

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}
}

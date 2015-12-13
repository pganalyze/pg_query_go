// Auto-generated - DO NOT EDIT

package pg_query

func (node LoadStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("LoadStmt")

	if node.Filename != nil {
		ctx.WriteString(*node.Filename)
	}
}

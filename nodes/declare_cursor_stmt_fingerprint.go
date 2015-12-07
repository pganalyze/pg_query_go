// Auto-generated - DO NOT EDIT

package pg_query

func (node DeclareCursorStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("DECLARECURSOR")

	if node.Portalname != nil {
		ctx.WriteString(*node.Portalname)
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}
}

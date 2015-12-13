// Auto-generated - DO NOT EDIT

package pg_query

func (node PrepareStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("PrepareStmt")

	for _, subNode := range node.Argtypes {
		subNode.Fingerprint(ctx, "Argtypes")
	}

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx, "Query")
	}
}

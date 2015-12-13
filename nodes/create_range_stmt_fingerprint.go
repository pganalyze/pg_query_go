// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateRangeStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateRangeStmt")

	for _, subNode := range node.Params {
		subNode.Fingerprint(ctx, "Params")
	}

	for _, subNode := range node.TypeName {
		subNode.Fingerprint(ctx, "TypeName")
	}
}

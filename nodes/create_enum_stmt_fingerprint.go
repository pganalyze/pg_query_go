// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateEnumStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateEnumStmt")

	for _, subNode := range node.TypeName {
		subNode.Fingerprint(ctx, "TypeName")
	}

	for _, subNode := range node.Vals {
		subNode.Fingerprint(ctx, "Vals")
	}
}

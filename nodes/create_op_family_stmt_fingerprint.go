// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateOpFamilyStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateOpFamilyStmt")

	if node.Amname != nil {
		ctx.WriteString(*node.Amname)
	}

	for _, subNode := range node.Opfamilyname {
		subNode.Fingerprint(ctx, "Opfamilyname")
	}
}

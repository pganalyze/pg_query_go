// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateOpFamilyStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateOpFamilyStmt")

	if node.Amname != nil {
		ctx.WriteString(*node.Amname)
	}

	node.Opfamilyname.Fingerprint(ctx, "Opfamilyname")
}

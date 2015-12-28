// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateOpFamilyStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateOpFamilyStmt")

	if node.Amname != nil {
		ctx.WriteString("amname")
		ctx.WriteString(*node.Amname)
	}

	if len(node.Opfamilyname.Items) > 0 {
		ctx.WriteString("opfamilyname")
		node.Opfamilyname.Fingerprint(ctx, "Opfamilyname")
	}
}

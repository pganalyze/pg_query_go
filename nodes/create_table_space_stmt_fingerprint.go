// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateTableSpaceStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateTableSpaceStmt")
	// Intentionally ignoring node.Location for fingerprinting

	node.Options.Fingerprint(ctx, "Options")

	if node.Owner != nil {
		ctx.WriteString(*node.Owner)
	}

	if node.Tablespacename != nil {
		ctx.WriteString(*node.Tablespacename)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateTableSpaceStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CreateTableSpaceStmt")
	// Intentionally ignoring node.Location for fingerprinting

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Owner != nil {
		ctx.WriteString(*node.Owner)
	}

	if node.Tablespacename != nil {
		ctx.WriteString(*node.Tablespacename)
	}
}

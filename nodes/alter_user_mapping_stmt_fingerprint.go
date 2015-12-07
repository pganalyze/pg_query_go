// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterUserMappingStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ALTERUSERMAPPINGSTMT")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Servername != nil {
		ctx.WriteString(*node.Servername)
	}

	if node.Username != nil {
		ctx.WriteString(*node.Username)
	}
}

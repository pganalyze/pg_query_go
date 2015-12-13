// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterTSDictionaryStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterTSDictionaryStmt")

	for _, subNode := range node.Dictname {
		subNode.Fingerprint(ctx, "Dictname")
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx, "Options")
	}
}

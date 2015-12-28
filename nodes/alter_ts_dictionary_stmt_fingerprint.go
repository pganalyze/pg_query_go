// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterTSDictionaryStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterTSDictionaryStmt")
	if len(node.Dictname.Items) > 0 {
		ctx.WriteString("dictname")
		node.Dictname.Fingerprint(ctx, "Dictname")
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}
}

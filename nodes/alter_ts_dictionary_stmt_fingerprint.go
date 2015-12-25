// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterTSDictionaryStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterTSDictionaryStmt")
	node.Dictname.Fingerprint(ctx, "Dictname")
	node.Options.Fingerprint(ctx, "Options")
}

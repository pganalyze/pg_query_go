// Auto-generated - DO NOT EDIT

package pg_query

func (node DeallocateStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("DEALLOCATESTMT")

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}
}

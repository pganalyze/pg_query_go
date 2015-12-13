// Auto-generated - DO NOT EDIT

package pg_query

func (node DeallocateStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("DeallocateStmt")

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}
}

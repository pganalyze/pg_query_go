// Auto-generated - DO NOT EDIT

package pg_query

func (node ColumnRef) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ColumnRef")

	for _, subNode := range node.Fields {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting
}

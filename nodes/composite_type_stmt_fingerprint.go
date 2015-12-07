// Auto-generated - DO NOT EDIT

package pg_query

func (node CompositeTypeStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("COMPOSITETYPESTMT")

	for _, subNode := range node.Coldeflist {
		subNode.Fingerprint(ctx)
	}

	if node.Typevar != nil {
		node.Typevar.Fingerprint(ctx)
	}
}

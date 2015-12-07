// Auto-generated - DO NOT EDIT

package pg_query

func (node LockRows) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("LOCKROWS")

	for _, subNode := range node.RowMarks {
		subNode.Fingerprint(ctx)
	}
}

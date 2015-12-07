// Auto-generated - DO NOT EDIT

package pg_query

func (node MergeAppend) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("MERGEAPPEND")

	for _, subNode := range node.Mergeplans {
		subNode.Fingerprint(ctx)
	}
}

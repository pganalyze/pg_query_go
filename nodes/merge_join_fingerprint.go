// Auto-generated - DO NOT EDIT

package pg_query

func (node MergeJoin) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("MERGEJOIN")

	for _, subNode := range node.Mergeclauses {
		subNode.Fingerprint(ctx)
	}
}

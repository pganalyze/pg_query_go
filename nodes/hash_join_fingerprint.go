// Auto-generated - DO NOT EDIT

package pg_query

func (node HashJoin) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("HASHJOIN")

	for _, subNode := range node.Hashclauses {
		subNode.Fingerprint(ctx)
	}
}

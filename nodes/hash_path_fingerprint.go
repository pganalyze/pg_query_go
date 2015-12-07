// Auto-generated - DO NOT EDIT

package pg_query

func (node HashPath) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("HASHPATH")

	for _, subNode := range node.PathHashclauses {
		subNode.Fingerprint(ctx)
	}
}

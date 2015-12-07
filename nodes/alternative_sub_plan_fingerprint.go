// Auto-generated - DO NOT EDIT

package pg_query

func (node AlternativeSubPlan) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ALTERNATIVESUBPLAN")

	for _, subNode := range node.Subplans {
		subNode.Fingerprint(ctx)
	}
}

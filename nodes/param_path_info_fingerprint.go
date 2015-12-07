// Auto-generated - DO NOT EDIT

package pg_query

func (node ParamPathInfo) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("PARAMPATHINFO")

	for _, subNode := range node.PpiClauses {
		subNode.Fingerprint(ctx)
	}
}

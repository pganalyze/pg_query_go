// Auto-generated - DO NOT EDIT

package pg_query

func (node AppendRelInfo) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("APPENDRELINFO")

	for _, subNode := range node.TranslatedVars {
		subNode.Fingerprint(ctx)
	}
}

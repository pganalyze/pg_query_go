// Auto-generated - DO NOT EDIT

package pg_query

func (node AppendPath) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("APPENDPATH")

	for _, subNode := range node.Subpaths {
		subNode.Fingerprint(ctx)
	}
}

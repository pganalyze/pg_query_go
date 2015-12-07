// Auto-generated - DO NOT EDIT

package pg_query

func (node MergeAppendPath) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("MERGEAPPENDPATH")

	for _, subNode := range node.Subpaths {
		subNode.Fingerprint(ctx)
	}
}

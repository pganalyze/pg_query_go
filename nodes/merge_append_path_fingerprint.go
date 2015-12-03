// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node MergeAppendPath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "MergeAppendPath")

	for _, subNode := range node.Subpaths {
		subNode.Fingerprint(ctx)
	}
}

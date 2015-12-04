// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AppendPath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "APPENDPATH")

	for _, subNode := range node.Subpaths {
		subNode.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ColumnRef) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ColumnRef")

	for _, subNode := range node.Fields {
		subNode.Fingerprint(ctx)
	}
}

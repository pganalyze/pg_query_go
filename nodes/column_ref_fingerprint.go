// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ColumnRef) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "COLUMNREF")

	for _, subNode := range node.Fields {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting
}

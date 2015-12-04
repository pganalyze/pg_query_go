// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CompositeTypeStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "COMPOSITETYPESTMT")

	for _, subNode := range node.Coldeflist {
		subNode.Fingerprint(ctx)
	}

	if node.Typevar != nil {
		node.Typevar.Fingerprint(ctx)
	}
}

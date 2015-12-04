// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateRangeStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATERANGESTMT")

	for _, subNode := range node.Params {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.TypeName {
		subNode.Fingerprint(ctx)
	}
}

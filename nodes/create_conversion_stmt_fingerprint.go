// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateConversionStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateConversionStmt")

	for _, subNode := range node.ConversionName {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.FuncName {
		subNode.Fingerprint(ctx)
	}
}

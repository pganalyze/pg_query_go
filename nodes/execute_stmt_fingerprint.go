// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ExecuteStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "EXECUTESTMT")

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}

	for _, subNode := range node.Params {
		subNode.Fingerprint(ctx)
	}
}

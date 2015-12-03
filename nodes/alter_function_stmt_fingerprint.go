// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterFunctionStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterFunctionStmt")

	for _, subNode := range node.Actions {
		subNode.Fingerprint(ctx)
	}

	if node.Func != nil {
		node.Func.Fingerprint(ctx)
	}
}

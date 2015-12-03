// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node VariableSetStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "VariableSetStmt")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}
}

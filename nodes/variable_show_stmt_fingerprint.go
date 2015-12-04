// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node VariableShowStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SHOW")

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}
}

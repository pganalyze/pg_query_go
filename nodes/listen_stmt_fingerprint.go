// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ListenStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "LISTENSTMT")

	if node.Conditionname != nil {
		io.WriteString(ctx.hash, *node.Conditionname)
	}
}

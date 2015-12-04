// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node UnlistenStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "UNLISTENSTMT")

	if node.Conditionname != nil {
		io.WriteString(ctx.hash, *node.Conditionname)
	}
}

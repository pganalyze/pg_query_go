// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node DeallocateStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DEALLOCATESTMT")

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}
}

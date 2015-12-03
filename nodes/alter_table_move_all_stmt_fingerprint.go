// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterTableMoveAllStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterTableMoveAllStmt")

	for _, subNode := range node.Roles {
		subNode.Fingerprint(ctx)
	}
}

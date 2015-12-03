// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterTableStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterTableStmt")

	for _, subNode := range node.Cmds {
		subNode.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}
}

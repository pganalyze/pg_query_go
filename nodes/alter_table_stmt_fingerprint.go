// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node AlterTableStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTER TABLE")

	for _, subNode := range node.Cmds {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.MissingOk))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Relkind)))
}

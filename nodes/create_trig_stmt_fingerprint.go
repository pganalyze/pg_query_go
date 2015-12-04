// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node CreateTrigStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATETRIGSTMT")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Columns {
		subNode.Fingerprint(ctx)
	}

	if node.Constrrel != nil {
		node.Constrrel.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Deferrable))

	for _, subNode := range node.Funcname {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Initdeferred))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Isconstraint))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Row))

	if node.Trigname != nil {
		io.WriteString(ctx.hash, *node.Trigname)
	}

	if node.WhenClause != nil {
		node.WhenClause.Fingerprint(ctx)
	}
}

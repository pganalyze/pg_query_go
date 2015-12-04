// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node RuleStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RULESTMT")

	for _, subNode := range node.Actions {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Event)))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Instead))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Replace))

	if node.Rulename != nil {
		io.WriteString(ctx.hash, *node.Rulename)
	}

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx)
	}
}

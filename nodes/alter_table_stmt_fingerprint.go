// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterTableStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("AlterTableStmt")

	for _, subNode := range node.Cmds {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.MissingOk))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Relkind)))
}

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterTableStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterTableStmt")

	for _, subNode := range node.Cmds {
		subNode.Fingerprint(ctx, "Cmds")
	}

	ctx.WriteString(strconv.FormatBool(node.MissingOk))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}

	ctx.WriteString(strconv.Itoa(int(node.Relkind)))
}

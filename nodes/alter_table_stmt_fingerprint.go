// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterTableStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterTableStmt")
	if len(node.Cmds.Items) > 0 {
		ctx.WriteString("cmds")
		node.Cmds.Fingerprint(ctx, "Cmds")
	}

	if node.MissingOk {
		ctx.WriteString("missing_ok")
		ctx.WriteString(strconv.FormatBool(node.MissingOk))
	}

	if node.Relation != nil {
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}

	if int(node.Relkind) != 0 {
		ctx.WriteString("relkind")
		ctx.WriteString(strconv.Itoa(int(node.Relkind)))
	}
}

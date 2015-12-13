// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterTableMoveAllStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("AlterTableMoveAllStmt")

	if node.NewTablespacename != nil {
		ctx.WriteString(*node.NewTablespacename)
	}

	ctx.WriteString(strconv.FormatBool(node.Nowait))
	ctx.WriteString(strconv.Itoa(int(node.Objtype)))

	if node.OrigTablespacename != nil {
		ctx.WriteString(*node.OrigTablespacename)
	}

	for _, subNode := range node.Roles {
		subNode.Fingerprint(ctx)
	}
}

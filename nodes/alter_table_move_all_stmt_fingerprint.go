// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterTableMoveAllStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterTableMoveAllStmt")

	if node.NewTablespacename != nil {
		ctx.WriteString("new_tablespacename")
		ctx.WriteString(*node.NewTablespacename)
	}

	if node.Nowait {
		ctx.WriteString("nowait")
		ctx.WriteString(strconv.FormatBool(node.Nowait))
	}

	if int(node.Objtype) != 0 {
		ctx.WriteString("objtype")
		ctx.WriteString(strconv.Itoa(int(node.Objtype)))
	}

	if node.OrigTablespacename != nil {
		ctx.WriteString("orig_tablespacename")
		ctx.WriteString(*node.OrigTablespacename)
	}

	if len(node.Roles.Items) > 0 {
		ctx.WriteString("roles")
		node.Roles.Fingerprint(ctx, "Roles")
	}
}

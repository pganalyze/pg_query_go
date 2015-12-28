// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateRoleStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateRoleStmt")
	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if node.Role != nil {
		ctx.WriteString("role")
		ctx.WriteString(*node.Role)
	}

	if int(node.StmtType) != 0 {
		ctx.WriteString("stmt_type")
		ctx.WriteString(strconv.Itoa(int(node.StmtType)))
	}
}

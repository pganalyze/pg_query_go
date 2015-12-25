// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateRoleStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateRoleStmt")
	node.Options.Fingerprint(ctx, "Options")

	if node.Role != nil {
		ctx.WriteString(*node.Role)
	}

	ctx.WriteString(strconv.Itoa(int(node.StmtType)))
}

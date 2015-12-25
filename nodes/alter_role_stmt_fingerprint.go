// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterRoleStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterRoleStmt")
	ctx.WriteString(strconv.Itoa(int(node.Action)))
	node.Options.Fingerprint(ctx, "Options")

	if node.Role != nil {
		ctx.WriteString(*node.Role)
	}
}

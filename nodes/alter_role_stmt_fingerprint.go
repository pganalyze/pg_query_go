// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterRoleStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterRoleStmt")

	if node.Action != 0 {
		ctx.WriteString("action")
		ctx.WriteString(strconv.Itoa(int(node.Action)))
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if node.Role != nil {
		ctx.WriteString("role")
		ctx.WriteString(*node.Role)
	}
}

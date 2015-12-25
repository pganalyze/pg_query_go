// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node GrantRoleStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("GrantRoleStmt")
	ctx.WriteString(strconv.FormatBool(node.AdminOpt))
	ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	node.GrantedRoles.Fingerprint(ctx, "GrantedRoles")
	node.GranteeRoles.Fingerprint(ctx, "GranteeRoles")

	if node.Grantor != nil {
		ctx.WriteString(*node.Grantor)
	}

	ctx.WriteString(strconv.FormatBool(node.IsGrant))
}

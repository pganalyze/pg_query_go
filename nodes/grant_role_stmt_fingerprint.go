// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node GrantRoleStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("GrantRoleStmt")
	ctx.WriteString(strconv.FormatBool(node.AdminOpt))
	ctx.WriteString(strconv.Itoa(int(node.Behavior)))

	for _, subNode := range node.GrantedRoles {
		subNode.Fingerprint(ctx, "GrantedRoles")
	}

	for _, subNode := range node.GranteeRoles {
		subNode.Fingerprint(ctx, "GranteeRoles")
	}

	if node.Grantor != nil {
		ctx.WriteString(*node.Grantor)
	}

	ctx.WriteString(strconv.FormatBool(node.IsGrant))
}

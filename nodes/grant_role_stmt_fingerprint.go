// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node GrantRoleStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("GRANTROLESTMT")
	ctx.WriteString(strconv.FormatBool(node.AdminOpt))
	ctx.WriteString(strconv.Itoa(int(node.Behavior)))

	for _, subNode := range node.GrantedRoles {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.GranteeRoles {
		subNode.Fingerprint(ctx)
	}

	if node.Grantor != nil {
		ctx.WriteString(*node.Grantor)
	}

	ctx.WriteString(strconv.FormatBool(node.IsGrant))
}

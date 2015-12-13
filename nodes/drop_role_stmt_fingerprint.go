// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DropRoleStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DropRoleStmt")
	ctx.WriteString(strconv.FormatBool(node.MissingOk))

	for _, subNode := range node.Roles {
		subNode.Fingerprint(ctx, "Roles")
	}
}

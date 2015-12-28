// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DropOwnedStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DropOwnedStmt")

	if int(node.Behavior) != 0 {
		ctx.WriteString("behavior")
		ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	}

	if len(node.Roles.Items) > 0 {
		ctx.WriteString("roles")
		node.Roles.Fingerprint(ctx, "Roles")
	}
}

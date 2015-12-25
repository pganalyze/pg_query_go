// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DropOwnedStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DropOwnedStmt")
	ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	node.Roles.Fingerprint(ctx, "Roles")
}

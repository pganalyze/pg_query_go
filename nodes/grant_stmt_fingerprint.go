// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node GrantStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("GrantStmt")
	ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	ctx.WriteString(strconv.FormatBool(node.GrantOption))
	node.Grantees.Fingerprint(ctx, "Grantees")
	ctx.WriteString(strconv.FormatBool(node.IsGrant))
	node.Objects.Fingerprint(ctx, "Objects")
	ctx.WriteString(strconv.Itoa(int(node.Objtype)))
	node.Privileges.Fingerprint(ctx, "Privileges")
	ctx.WriteString(strconv.Itoa(int(node.Targtype)))
}

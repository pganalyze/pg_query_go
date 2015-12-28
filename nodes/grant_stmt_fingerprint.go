// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node GrantStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("GrantStmt")

	if int(node.Behavior) != 0 {
		ctx.WriteString("behavior")
		ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	}

	if node.GrantOption {
		ctx.WriteString("grant_option")
		ctx.WriteString(strconv.FormatBool(node.GrantOption))
	}

	if len(node.Grantees.Items) > 0 {
		ctx.WriteString("grantees")
		node.Grantees.Fingerprint(ctx, "Grantees")
	}

	if node.IsGrant {
		ctx.WriteString("is_grant")
		ctx.WriteString(strconv.FormatBool(node.IsGrant))
	}

	if len(node.Objects.Items) > 0 {
		ctx.WriteString("objects")
		node.Objects.Fingerprint(ctx, "Objects")
	}

	if int(node.Objtype) != 0 {
		ctx.WriteString("objtype")
		ctx.WriteString(strconv.Itoa(int(node.Objtype)))
	}

	if len(node.Privileges.Items) > 0 {
		ctx.WriteString("privileges")
		node.Privileges.Fingerprint(ctx, "Privileges")
	}

	if int(node.Targtype) != 0 {
		ctx.WriteString("targtype")
		ctx.WriteString(strconv.Itoa(int(node.Targtype)))
	}
}

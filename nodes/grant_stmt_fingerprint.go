// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node GrantStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("GrantStmt")
	ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	ctx.WriteString(strconv.FormatBool(node.GrantOption))

	for _, subNode := range node.Grantees {
		subNode.Fingerprint(ctx, "Grantees")
	}

	ctx.WriteString(strconv.FormatBool(node.IsGrant))

	for _, subNode := range node.Objects {
		subNode.Fingerprint(ctx, "Objects")
	}

	ctx.WriteString(strconv.Itoa(int(node.Objtype)))

	for _, subNode := range node.Privileges {
		subNode.Fingerprint(ctx, "Privileges")
	}

	ctx.WriteString(strconv.Itoa(int(node.Targtype)))
}

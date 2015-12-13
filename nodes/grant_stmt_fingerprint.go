// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node GrantStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("GrantStmt")
	ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	ctx.WriteString(strconv.FormatBool(node.GrantOption))

	for _, subNode := range node.Grantees {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.IsGrant))

	for _, subNode := range node.Objects {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Objtype)))

	for _, subNode := range node.Privileges {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Targtype)))
}

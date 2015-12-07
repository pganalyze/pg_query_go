// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DropUserMappingStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("DROPUSERMAPPINGSTMT")
	ctx.WriteString(strconv.FormatBool(node.MissingOk))

	if node.Servername != nil {
		ctx.WriteString(*node.Servername)
	}

	if node.Username != nil {
		ctx.WriteString(*node.Username)
	}
}

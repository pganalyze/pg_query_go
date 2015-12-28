// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DropUserMappingStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DropUserMappingStmt")

	if node.MissingOk {
		ctx.WriteString("missing_ok")
		ctx.WriteString(strconv.FormatBool(node.MissingOk))
	}

	if node.Servername != nil {
		ctx.WriteString("servername")
		ctx.WriteString(*node.Servername)
	}

	if node.Username != nil {
		ctx.WriteString("username")
		ctx.WriteString(*node.Username)
	}
}

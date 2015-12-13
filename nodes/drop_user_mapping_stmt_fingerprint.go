// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DropUserMappingStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DropUserMappingStmt")
	ctx.WriteString(strconv.FormatBool(node.MissingOk))

	if node.Servername != nil {
		ctx.WriteString(*node.Servername)
	}

	if node.Username != nil {
		ctx.WriteString(*node.Username)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateExtensionStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateExtensionStmt")

	if node.Extname != nil {
		ctx.WriteString(*node.Extname)
	}

	ctx.WriteString(strconv.FormatBool(node.IfNotExists))
	node.Options.Fingerprint(ctx, "Options")
}

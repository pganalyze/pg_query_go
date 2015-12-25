// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateSchemaStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateSchemaStmt")

	if node.Authid != nil {
		ctx.WriteString(*node.Authid)
	}

	ctx.WriteString(strconv.FormatBool(node.IfNotExists))
	node.SchemaElts.Fingerprint(ctx, "SchemaElts")

	if node.Schemaname != nil {
		ctx.WriteString(*node.Schemaname)
	}
}

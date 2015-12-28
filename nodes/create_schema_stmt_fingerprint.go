// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateSchemaStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateSchemaStmt")

	if node.Authid != nil {
		ctx.WriteString("authid")
		ctx.WriteString(*node.Authid)
	}

	if node.IfNotExists {
		ctx.WriteString("if_not_exists")
		ctx.WriteString(strconv.FormatBool(node.IfNotExists))
	}

	if len(node.SchemaElts.Items) > 0 {
		ctx.WriteString("schemaElts")
		node.SchemaElts.Fingerprint(ctx, "SchemaElts")
	}

	if node.Schemaname != nil {
		ctx.WriteString("schemaname")
		ctx.WriteString(*node.Schemaname)
	}
}

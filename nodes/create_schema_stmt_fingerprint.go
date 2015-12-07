// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateSchemaStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CREATE SCHEMA")

	if node.Authid != nil {
		ctx.WriteString(*node.Authid)
	}

	ctx.WriteString(strconv.FormatBool(node.IfNotExists))

	for _, subNode := range node.SchemaElts {
		subNode.Fingerprint(ctx)
	}

	if node.Schemaname != nil {
		ctx.WriteString(*node.Schemaname)
	}
}

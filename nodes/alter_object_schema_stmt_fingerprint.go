// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterObjectSchemaStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterObjectSchemaStmt")
	ctx.WriteString(strconv.FormatBool(node.MissingOk))

	if node.Newschema != nil {
		ctx.WriteString(*node.Newschema)
	}

	node.Objarg.Fingerprint(ctx, "Objarg")
	node.Object.Fingerprint(ctx, "Object")
	ctx.WriteString(strconv.Itoa(int(node.ObjectType)))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterObjectSchemaStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterObjectSchemaStmt")

	if node.MissingOk {
		ctx.WriteString("missing_ok")
		ctx.WriteString(strconv.FormatBool(node.MissingOk))
	}

	if node.Newschema != nil {
		ctx.WriteString("newschema")
		ctx.WriteString(*node.Newschema)
	}

	if len(node.Objarg.Items) > 0 {
		ctx.WriteString("objarg")
		node.Objarg.Fingerprint(ctx, "Objarg")
	}

	if len(node.Object.Items) > 0 {
		ctx.WriteString("object")
		node.Object.Fingerprint(ctx, "Object")
	}

	if int(node.ObjectType) != 0 {
		ctx.WriteString("objectType")
		ctx.WriteString(strconv.Itoa(int(node.ObjectType)))
	}

	if node.Relation != nil {
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}
}

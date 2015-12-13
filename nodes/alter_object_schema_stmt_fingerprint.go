// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterObjectSchemaStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterObjectSchemaStmt")
	ctx.WriteString(strconv.FormatBool(node.MissingOk))

	if node.Newschema != nil {
		ctx.WriteString(*node.Newschema)
	}

	for _, subNode := range node.Objarg {
		subNode.Fingerprint(ctx, "Objarg")
	}

	for _, subNode := range node.Object {
		subNode.Fingerprint(ctx, "Object")
	}

	ctx.WriteString(strconv.Itoa(int(node.ObjectType)))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}
}

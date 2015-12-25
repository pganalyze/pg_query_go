// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterOwnerStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterOwnerStmt")

	if node.Newowner != nil {
		ctx.WriteString(*node.Newowner)
	}

	node.Objarg.Fingerprint(ctx, "Objarg")
	node.Object.Fingerprint(ctx, "Object")
	ctx.WriteString(strconv.Itoa(int(node.ObjectType)))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}
}

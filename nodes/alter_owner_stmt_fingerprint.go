// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterOwnerStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterOwnerStmt")

	if node.Newowner != nil {
		ctx.WriteString("newowner")
		ctx.WriteString(*node.Newowner)
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

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RenameStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RenameStmt")
	ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	ctx.WriteString(strconv.FormatBool(node.MissingOk))

	if node.Newname != nil {
		ctx.WriteString(*node.Newname)
	}

	node.Objarg.Fingerprint(ctx, "Objarg")
	node.Object.Fingerprint(ctx, "Object")

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}

	ctx.WriteString(strconv.Itoa(int(node.RelationType)))
	ctx.WriteString(strconv.Itoa(int(node.RenameType)))

	if node.Subname != nil {
		ctx.WriteString(*node.Subname)
	}
}

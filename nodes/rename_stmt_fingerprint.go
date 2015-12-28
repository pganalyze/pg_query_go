// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RenameStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RenameStmt")

	if int(node.Behavior) != 0 {
		ctx.WriteString("behavior")
		ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	}

	if node.MissingOk {
		ctx.WriteString("missing_ok")
		ctx.WriteString(strconv.FormatBool(node.MissingOk))
	}

	if node.Newname != nil {
		ctx.WriteString("newname")
		ctx.WriteString(*node.Newname)
	}

	if len(node.Objarg.Items) > 0 {
		ctx.WriteString("objarg")
		node.Objarg.Fingerprint(ctx, "Objarg")
	}

	if len(node.Object.Items) > 0 {
		ctx.WriteString("object")
		node.Object.Fingerprint(ctx, "Object")
	}

	if node.Relation != nil {
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}

	if int(node.RelationType) != 0 {
		ctx.WriteString("relationType")
		ctx.WriteString(strconv.Itoa(int(node.RelationType)))
	}

	if int(node.RenameType) != 0 {
		ctx.WriteString("renameType")
		ctx.WriteString(strconv.Itoa(int(node.RenameType)))
	}

	if node.Subname != nil {
		ctx.WriteString("subname")
		ctx.WriteString(*node.Subname)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CommentStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CommentStmt")

	if node.Comment != nil {
		ctx.WriteString("comment")
		ctx.WriteString(*node.Comment)
	}

	if len(node.Objargs.Items) > 0 {
		ctx.WriteString("objargs")
		node.Objargs.Fingerprint(ctx, "Objargs")
	}

	if len(node.Objname.Items) > 0 {
		ctx.WriteString("objname")
		node.Objname.Fingerprint(ctx, "Objname")
	}

	if int(node.Objtype) != 0 {
		ctx.WriteString("objtype")
		ctx.WriteString(strconv.Itoa(int(node.Objtype)))
	}
}

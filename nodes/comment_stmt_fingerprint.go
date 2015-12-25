// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CommentStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CommentStmt")

	if node.Comment != nil {
		ctx.WriteString(*node.Comment)
	}

	node.Objargs.Fingerprint(ctx, "Objargs")
	node.Objname.Fingerprint(ctx, "Objname")
	ctx.WriteString(strconv.Itoa(int(node.Objtype)))
}

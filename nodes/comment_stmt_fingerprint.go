// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CommentStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CommentStmt")

	if node.Comment != nil {
		ctx.WriteString(*node.Comment)
	}

	for _, subNode := range node.Objargs {
		subNode.Fingerprint(ctx, "Objargs")
	}

	for _, subNode := range node.Objname {
		subNode.Fingerprint(ctx, "Objname")
	}

	ctx.WriteString(strconv.Itoa(int(node.Objtype)))
}

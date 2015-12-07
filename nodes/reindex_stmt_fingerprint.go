// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ReindexStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("REINDEXSTMT")
	ctx.WriteString(strconv.FormatBool(node.DoSystem))
	ctx.WriteString(strconv.FormatBool(node.DoUser))
	ctx.WriteString(strconv.Itoa(int(node.Kind)))

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}
}

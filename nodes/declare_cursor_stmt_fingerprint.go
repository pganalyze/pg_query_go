// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DeclareCursorStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("DeclareCursorStmt")
	ctx.WriteString(strconv.Itoa(int(node.Options)))

	if node.Portalname != nil {
		ctx.WriteString(*node.Portalname)
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}
}

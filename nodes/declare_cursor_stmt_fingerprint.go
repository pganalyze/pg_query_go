// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DeclareCursorStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DeclareCursorStmt")

	if node.Options != 0 {
		ctx.WriteString("options")
		ctx.WriteString(strconv.Itoa(int(node.Options)))
	}

	if node.Portalname != nil {
		ctx.WriteString("portalname")
		ctx.WriteString(*node.Portalname)
	}

	if node.Query != nil {
		ctx.WriteString("query")
		node.Query.Fingerprint(ctx, "Query")
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateTableAsStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateTableAsStmt")

	if node.Into != nil {
		node.Into.Fingerprint(ctx, "Into")
	}

	ctx.WriteString(strconv.FormatBool(node.IsSelectInto))

	if node.Query != nil {
		node.Query.Fingerprint(ctx, "Query")
	}

	ctx.WriteString(strconv.Itoa(int(node.Relkind)))
}

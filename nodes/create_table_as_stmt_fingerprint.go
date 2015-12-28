// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateTableAsStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateTableAsStmt")

	if node.Into != nil {
		ctx.WriteString("into")
		node.Into.Fingerprint(ctx, "Into")
	}

	if node.IsSelectInto {
		ctx.WriteString("is_select_into")
		ctx.WriteString(strconv.FormatBool(node.IsSelectInto))
	}

	if node.Query != nil {
		ctx.WriteString("query")
		node.Query.Fingerprint(ctx, "Query")
	}

	if int(node.Relkind) != 0 {
		ctx.WriteString("relkind")
		ctx.WriteString(strconv.Itoa(int(node.Relkind)))
	}
}

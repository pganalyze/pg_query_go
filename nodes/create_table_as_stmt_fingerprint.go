// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateTableAsStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CREATE TABLE AS")

	if node.Into != nil {
		node.Into.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.IsSelectInto))

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Relkind)))
}

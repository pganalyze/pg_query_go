// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node CreateTableAsStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATE TABLE AS")

	if node.Into != nil {
		node.Into.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.IsSelectInto))

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Relkind)))
}

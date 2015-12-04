// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node CreateSchemaStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATE SCHEMA")

	if node.Authid != nil {
		io.WriteString(ctx.hash, *node.Authid)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.IfNotExists))

	for _, subNode := range node.SchemaElts {
		subNode.Fingerprint(ctx)
	}

	if node.Schemaname != nil {
		io.WriteString(ctx.hash, *node.Schemaname)
	}
}

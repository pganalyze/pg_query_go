// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateSchemaStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateSchemaStmt")

	for _, subNode := range node.SchemaElts {
		subNode.Fingerprint(ctx)
	}
}

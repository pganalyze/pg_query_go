// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterEnumStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterEnumStmt")

	for _, subNode := range node.TypeName {
		subNode.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateTableSpaceStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateTableSpaceStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}

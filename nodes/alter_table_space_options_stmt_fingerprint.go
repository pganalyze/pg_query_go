// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterTableSpaceOptionsStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterTableSpaceOptionsStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterUserMappingStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterUserMappingStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}

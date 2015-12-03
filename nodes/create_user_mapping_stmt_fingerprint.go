// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateUserMappingStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateUserMappingStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}

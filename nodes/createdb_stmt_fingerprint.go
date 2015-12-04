// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreatedbStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATEDBSTMT")

	if node.Dbname != nil {
		io.WriteString(ctx.hash, *node.Dbname)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}

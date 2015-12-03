// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CommentStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CommentStmt")

	for _, subNode := range node.Objargs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Objname {
		subNode.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node CommentStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "COMMENTSTMT")

	if node.Comment != nil {
		io.WriteString(ctx.hash, *node.Comment)
	}

	for _, subNode := range node.Objargs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Objname {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Objtype)))
}

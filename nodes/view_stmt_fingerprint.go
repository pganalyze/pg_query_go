// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node ViewStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "VIEWSTMT")

	for _, subNode := range node.Aliases {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Replace))

	if node.View != nil {
		node.View.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.WithCheckOption)))
}

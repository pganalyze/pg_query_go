// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node CopyStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "COPY")

	for _, subNode := range node.Attlist {
		subNode.Fingerprint(ctx)
	}

	if node.Filename != nil {
		io.WriteString(ctx.hash, *node.Filename)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.IsFrom))
	io.WriteString(ctx.hash, strconv.FormatBool(node.IsProgram))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node AlterForeignServerStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTERFOREIGNSERVERSTMT")
	io.WriteString(ctx.hash, strconv.FormatBool(node.HasVersion))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Servername != nil {
		io.WriteString(ctx.hash, *node.Servername)
	}

	if node.Version != nil {
		io.WriteString(ctx.hash, *node.Version)
	}
}

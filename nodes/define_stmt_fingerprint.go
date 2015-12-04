// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node DefineStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DEFINESTMT")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Definition {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Defnames {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Kind)))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Oldstyle))
}

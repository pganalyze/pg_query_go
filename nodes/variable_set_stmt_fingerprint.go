// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node VariableSetStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SET")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.IsLocal))
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Kind)))

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}
}

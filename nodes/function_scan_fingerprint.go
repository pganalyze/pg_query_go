// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node FunctionScan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FUNCTIONSCAN")
	io.WriteString(ctx.hash, strconv.FormatBool(node.Funcordinality))

	for _, subNode := range node.Functions {
		subNode.Fingerprint(ctx)
	}
}

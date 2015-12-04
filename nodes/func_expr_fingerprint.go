// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node FuncExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FUNCEXPR")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Funcformat)))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Funcretset))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Funcvariadic))
	// Intentionally ignoring node.Location for fingerprinting
}

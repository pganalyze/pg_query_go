// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node FunctionParameter) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FUNCTIONPARAMETER")

	if node.ArgType != nil {
		node.ArgType.Fingerprint(ctx)
	}

	if node.Defexpr != nil {
		node.Defexpr.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Mode)))

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}
}

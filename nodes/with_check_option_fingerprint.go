// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node WithCheckOption) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "WITHCHECKOPTION")
	io.WriteString(ctx.hash, strconv.FormatBool(node.Cascaded))

	if node.Qual != nil {
		node.Qual.Fingerprint(ctx)
	}

	if node.Viewname != nil {
		io.WriteString(ctx.hash, *node.Viewname)
	}
}

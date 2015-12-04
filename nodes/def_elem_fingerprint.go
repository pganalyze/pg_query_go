// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node DefElem) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DEFELEM")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Defaction)))

	if node.Defname != nil {
		io.WriteString(ctx.hash, *node.Defname)
	}

	if node.Defnamespace != nil {
		io.WriteString(ctx.hash, *node.Defnamespace)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node RangeFunction) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RANGEFUNCTION")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx)
	}

	for _, subNode := range node.Coldeflist {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Functions {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.IsRowsfrom))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Lateral))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Ordinality))
}

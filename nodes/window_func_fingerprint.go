// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node WindowFunc) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "WINDOWFUNC")

	if node.Aggfilter != nil {
		node.Aggfilter.Fingerprint(ctx)
	}

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	io.WriteString(ctx.hash, strconv.FormatBool(node.Winagg))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Winstar))
}

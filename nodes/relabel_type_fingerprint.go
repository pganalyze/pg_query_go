// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node RelabelType) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RELABELTYPE")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Relabelformat)))
}

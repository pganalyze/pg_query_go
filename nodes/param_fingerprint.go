// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node Param) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PARAM")
	// Intentionally ignoring node.Location for fingerprinting

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Paramkind)))
}

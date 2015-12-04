// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node PathKey) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PATHKEY")

	if node.PkEclass != nil {
		node.PkEclass.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.PkNullsFirst))
}

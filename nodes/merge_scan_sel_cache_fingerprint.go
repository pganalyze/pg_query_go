// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node MergeScanSelCache) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "MERGESCANSELCACHE")
	io.WriteString(ctx.hash, strconv.FormatBool(node.NullsFirst))
}

// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node RangeSubselect) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RANGESUBSELECT")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Lateral))

	if node.Subquery != nil {
		node.Subquery.Fingerprint(ctx)
	}
}

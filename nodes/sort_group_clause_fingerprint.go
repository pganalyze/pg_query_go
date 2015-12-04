// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node SortGroupClause) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SORTGROUPCLAUSE")
	io.WriteString(ctx.hash, strconv.FormatBool(node.Hashable))
	io.WriteString(ctx.hash, strconv.FormatBool(node.NullsFirst))
}

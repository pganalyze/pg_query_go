// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node Agg) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AGG")
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Aggstrategy)))
}

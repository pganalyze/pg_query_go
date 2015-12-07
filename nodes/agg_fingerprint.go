// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Agg) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("AGG")
	ctx.WriteString(strconv.Itoa(int(node.Aggstrategy)))
}

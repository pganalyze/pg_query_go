// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SortGroupClause) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("SORTGROUPCLAUSE")
	ctx.WriteString(strconv.FormatBool(node.Hashable))
	ctx.WriteString(strconv.FormatBool(node.NullsFirst))
}

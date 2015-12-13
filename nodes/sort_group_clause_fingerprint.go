// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SortGroupClause) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("SortGroupClause")
	ctx.WriteString(strconv.Itoa(int(node.Eqop)))
	ctx.WriteString(strconv.FormatBool(node.Hashable))
	ctx.WriteString(strconv.FormatBool(node.NullsFirst))
	ctx.WriteString(strconv.Itoa(int(node.Sortop)))
	ctx.WriteString(strconv.Itoa(int(node.TleSortGroupRef)))
}

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RowMarkClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RowMarkClause")
	ctx.WriteString(strconv.FormatBool(node.NoWait))
	ctx.WriteString(strconv.FormatBool(node.PushedDown))
	ctx.WriteString(strconv.Itoa(int(node.Rti)))
	ctx.WriteString(strconv.Itoa(int(node.Strength)))
}

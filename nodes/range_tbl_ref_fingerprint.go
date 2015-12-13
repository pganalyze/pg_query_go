// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeTblRef) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RangeTblRef")
	ctx.WriteString(strconv.Itoa(int(node.Rtindex)))
}

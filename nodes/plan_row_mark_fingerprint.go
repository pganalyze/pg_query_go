// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node PlanRowMark) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("PLANROWMARK")
	ctx.WriteString(strconv.FormatBool(node.IsParent))
	ctx.WriteString(strconv.Itoa(int(node.MarkType)))
	ctx.WriteString(strconv.FormatBool(node.NoWait))
}

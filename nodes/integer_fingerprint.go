// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Integer) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("Integer")

	if node.Ival != 0 {
		ctx.WriteString("ival")
		ctx.WriteString(strconv.Itoa(int(node.Ival)))
	}
}

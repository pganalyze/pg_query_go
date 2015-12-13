// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Integer) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("Integer")
	ctx.WriteString(strconv.Itoa(int(node.Ival)))
}

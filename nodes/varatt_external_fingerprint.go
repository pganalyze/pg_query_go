// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node varatt_external) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("varatt_external")
	ctx.WriteString(strconv.Itoa(int(node.VaExtsize)))
	ctx.WriteString(strconv.Itoa(int(node.VaRawsize)))
	ctx.WriteString(strconv.Itoa(int(node.VaToastrelid)))
	ctx.WriteString(strconv.Itoa(int(node.VaValueid)))
}

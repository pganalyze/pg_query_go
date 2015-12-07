// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Const) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CONST")
	ctx.WriteString(strconv.FormatBool(node.Constbyval))
	ctx.WriteString(strconv.FormatBool(node.Constisnull))
	// Intentionally ignoring node.Location for fingerprinting
}

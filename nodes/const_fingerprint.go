// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Const) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("Const")
	ctx.WriteString(strconv.FormatBool(node.Constbyval))
	ctx.WriteString(strconv.Itoa(int(node.Constcollid)))
	ctx.WriteString(strconv.FormatBool(node.Constisnull))
	ctx.WriteString(strconv.Itoa(int(node.Constlen)))
	ctx.WriteString(strconv.Itoa(int(node.Consttype)))
	ctx.WriteString(strconv.Itoa(int(node.Consttypmod)))
	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

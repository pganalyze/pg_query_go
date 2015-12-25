// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node MinMaxExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("MinMaxExpr")
	node.Args.Fingerprint(ctx, "Args")
	ctx.WriteString(strconv.Itoa(int(node.Inputcollid)))
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Minmaxcollid)))
	ctx.WriteString(strconv.Itoa(int(node.Minmaxtype)))
	ctx.WriteString(strconv.Itoa(int(node.Op)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

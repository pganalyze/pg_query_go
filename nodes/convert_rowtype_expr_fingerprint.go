// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ConvertRowtypeExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ConvertRowtypeExpr")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx, "Arg")
	}

	ctx.WriteString(strconv.Itoa(int(node.Convertformat)))
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Resulttype)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

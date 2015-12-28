// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ConvertRowtypeExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ConvertRowtypeExpr")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	if int(node.Convertformat) != 0 {
		ctx.WriteString("convertformat")
		ctx.WriteString(strconv.Itoa(int(node.Convertformat)))
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Resulttype != 0 {
		ctx.WriteString("resulttype")
		ctx.WriteString(strconv.Itoa(int(node.Resulttype)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

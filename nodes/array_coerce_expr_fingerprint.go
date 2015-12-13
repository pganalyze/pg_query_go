// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ArrayCoerceExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ArrayCoerceExpr")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx, "Arg")
	}

	ctx.WriteString(strconv.Itoa(int(node.Coerceformat)))
	ctx.WriteString(strconv.Itoa(int(node.Elemfuncid)))
	ctx.WriteString(strconv.FormatBool(node.IsExplicit))
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Resultcollid)))
	ctx.WriteString(strconv.Itoa(int(node.Resulttype)))
	ctx.WriteString(strconv.Itoa(int(node.Resulttypmod)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

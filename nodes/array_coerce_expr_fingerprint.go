// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ArrayCoerceExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ArrayCoerceExpr")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	if int(node.Coerceformat) != 0 {
		ctx.WriteString("coerceformat")
		ctx.WriteString(strconv.Itoa(int(node.Coerceformat)))
	}

	if node.Elemfuncid != 0 {
		ctx.WriteString("elemfuncid")
		ctx.WriteString(strconv.Itoa(int(node.Elemfuncid)))
	}

	if node.IsExplicit {
		ctx.WriteString("isExplicit")
		ctx.WriteString(strconv.FormatBool(node.IsExplicit))
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Resultcollid != 0 {
		ctx.WriteString("resultcollid")
		ctx.WriteString(strconv.Itoa(int(node.Resultcollid)))
	}

	if node.Resulttype != 0 {
		ctx.WriteString("resulttype")
		ctx.WriteString(strconv.Itoa(int(node.Resulttype)))
	}

	if node.Resulttypmod != 0 {
		ctx.WriteString("resulttypmod")
		ctx.WriteString(strconv.Itoa(int(node.Resulttypmod)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

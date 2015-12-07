// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ArrayCoerceExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ARRAYCOERCEEXPR")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Coerceformat)))
	ctx.WriteString(strconv.FormatBool(node.IsExplicit))
	// Intentionally ignoring node.Location for fingerprinting
}

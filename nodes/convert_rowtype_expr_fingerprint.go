// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ConvertRowtypeExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ConvertRowtypeExpr")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Convertformat)))
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Resulttype)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CollateExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CollateExpr")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx, "Arg")
	}

	ctx.WriteString(strconv.Itoa(int(node.CollOid)))
	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

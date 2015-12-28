// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CollateExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CollateExpr")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	if node.CollOid != 0 {
		ctx.WriteString("collOid")
		ctx.WriteString(strconv.Itoa(int(node.CollOid)))
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

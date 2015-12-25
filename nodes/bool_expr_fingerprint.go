// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node BoolExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("BoolExpr")
	node.Args.Fingerprint(ctx, "Args")
	ctx.WriteString(strconv.Itoa(int(node.Boolop)))
	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

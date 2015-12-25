// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CoalesceExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CoalesceExpr")
	node.Args.Fingerprint(ctx, "Args")
	ctx.WriteString(strconv.Itoa(int(node.Coalescecollid)))
	ctx.WriteString(strconv.Itoa(int(node.Coalescetype)))
	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

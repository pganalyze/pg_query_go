// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CoalesceExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CoalesceExpr")
	if len(node.Args.Items) > 0 {
		ctx.WriteString("args")
		node.Args.Fingerprint(ctx, "Args")
	}

	if node.Coalescecollid != 0 {
		ctx.WriteString("coalescecollid")
		ctx.WriteString(strconv.Itoa(int(node.Coalescecollid)))
	}

	if node.Coalescetype != 0 {
		ctx.WriteString("coalescetype")
		ctx.WriteString(strconv.Itoa(int(node.Coalescetype)))
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

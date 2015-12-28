// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node MinMaxExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("MinMaxExpr")
	if len(node.Args.Items) > 0 {
		ctx.WriteString("args")
		node.Args.Fingerprint(ctx, "Args")
	}

	if node.Inputcollid != 0 {
		ctx.WriteString("inputcollid")
		ctx.WriteString(strconv.Itoa(int(node.Inputcollid)))
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Minmaxcollid != 0 {
		ctx.WriteString("minmaxcollid")
		ctx.WriteString(strconv.Itoa(int(node.Minmaxcollid)))
	}

	if node.Minmaxtype != 0 {
		ctx.WriteString("minmaxtype")
		ctx.WriteString(strconv.Itoa(int(node.Minmaxtype)))
	}

	if int(node.Op) != 0 {
		ctx.WriteString("op")
		ctx.WriteString(strconv.Itoa(int(node.Op)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

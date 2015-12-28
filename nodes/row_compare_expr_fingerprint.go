// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RowCompareExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RowCompareExpr")
	if len(node.Inputcollids.Items) > 0 {
		ctx.WriteString("inputcollids")
		node.Inputcollids.Fingerprint(ctx, "Inputcollids")
	}

	if len(node.Largs.Items) > 0 {
		ctx.WriteString("largs")
		node.Largs.Fingerprint(ctx, "Largs")
	}

	if len(node.Opfamilies.Items) > 0 {
		ctx.WriteString("opfamilies")
		node.Opfamilies.Fingerprint(ctx, "Opfamilies")
	}

	if len(node.Opnos.Items) > 0 {
		ctx.WriteString("opnos")
		node.Opnos.Fingerprint(ctx, "Opnos")
	}

	if len(node.Rargs.Items) > 0 {
		ctx.WriteString("rargs")
		node.Rargs.Fingerprint(ctx, "Rargs")
	}

	if int(node.Rctype) != 0 {
		ctx.WriteString("rctype")
		ctx.WriteString(strconv.Itoa(int(node.Rctype)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

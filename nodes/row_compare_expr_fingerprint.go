// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RowCompareExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RowCompareExpr")
	node.Inputcollids.Fingerprint(ctx, "Inputcollids")
	node.Largs.Fingerprint(ctx, "Largs")
	node.Opfamilies.Fingerprint(ctx, "Opfamilies")
	node.Opnos.Fingerprint(ctx, "Opnos")
	node.Rargs.Fingerprint(ctx, "Rargs")
	ctx.WriteString(strconv.Itoa(int(node.Rctype)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

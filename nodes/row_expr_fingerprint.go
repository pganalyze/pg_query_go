// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RowExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RowExpr")
	node.Args.Fingerprint(ctx, "Args")
	node.Colnames.Fingerprint(ctx, "Colnames")
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.RowFormat)))
	ctx.WriteString(strconv.Itoa(int(node.RowTypeid)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

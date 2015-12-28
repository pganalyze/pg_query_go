// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RowExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RowExpr")
	if len(node.Args.Items) > 0 {
		ctx.WriteString("args")
		node.Args.Fingerprint(ctx, "Args")
	}

	if len(node.Colnames.Items) > 0 {
		ctx.WriteString("colnames")
		node.Colnames.Fingerprint(ctx, "Colnames")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if int(node.RowFormat) != 0 {
		ctx.WriteString("row_format")
		ctx.WriteString(strconv.Itoa(int(node.RowFormat)))
	}

	if node.RowTypeid != 0 {
		ctx.WriteString("row_typeid")
		ctx.WriteString(strconv.Itoa(int(node.RowTypeid)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

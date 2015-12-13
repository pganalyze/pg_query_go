// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TargetEntry) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("TargetEntry")

	if node.Expr != nil {
		node.Expr.Fingerprint(ctx, "Expr")
	}

	ctx.WriteString(strconv.FormatBool(node.Resjunk))

	if node.Resname != nil {
		ctx.WriteString(*node.Resname)
	}

	ctx.WriteString(strconv.Itoa(int(node.Resno)))
	ctx.WriteString(strconv.Itoa(int(node.Resorigcol)))
	ctx.WriteString(strconv.Itoa(int(node.Resorigtbl)))
	ctx.WriteString(strconv.Itoa(int(node.Ressortgroupref)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

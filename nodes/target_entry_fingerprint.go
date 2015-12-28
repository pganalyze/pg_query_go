// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TargetEntry) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("TargetEntry")

	if node.Expr != nil {
		ctx.WriteString("expr")
		node.Expr.Fingerprint(ctx, "Expr")
	}

	if node.Resjunk {
		ctx.WriteString("resjunk")
		ctx.WriteString(strconv.FormatBool(node.Resjunk))
	}

	if node.Resname != nil {
		ctx.WriteString("resname")
		ctx.WriteString(*node.Resname)
	}

	if node.Resno != 0 {
		ctx.WriteString("resno")
		ctx.WriteString(strconv.Itoa(int(node.Resno)))
	}

	if node.Resorigcol != 0 {
		ctx.WriteString("resorigcol")
		ctx.WriteString(strconv.Itoa(int(node.Resorigcol)))
	}

	if node.Resorigtbl != 0 {
		ctx.WriteString("resorigtbl")
		ctx.WriteString(strconv.Itoa(int(node.Resorigtbl)))
	}

	if node.Ressortgroupref != 0 {
		ctx.WriteString("ressortgroupref")
		ctx.WriteString(strconv.Itoa(int(node.Ressortgroupref)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

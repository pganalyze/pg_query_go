// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node JoinExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("JoinExpr")

	if node.Alias != nil {
		ctx.WriteString("alias")
		node.Alias.Fingerprint(ctx, "Alias")
	}

	if node.IsNatural {
		ctx.WriteString("isNatural")
		ctx.WriteString(strconv.FormatBool(node.IsNatural))
	}

	if int(node.Jointype) != 0 {
		ctx.WriteString("jointype")
		ctx.WriteString(strconv.Itoa(int(node.Jointype)))
	}

	if node.Larg != nil {
		ctx.WriteString("larg")
		node.Larg.Fingerprint(ctx, "Larg")
	}

	if node.Quals != nil {
		ctx.WriteString("quals")
		node.Quals.Fingerprint(ctx, "Quals")
	}

	if node.Rarg != nil {
		ctx.WriteString("rarg")
		node.Rarg.Fingerprint(ctx, "Rarg")
	}

	if node.Rtindex != 0 {
		ctx.WriteString("rtindex")
		ctx.WriteString(strconv.Itoa(int(node.Rtindex)))
	}

	if len(node.UsingClause.Items) > 0 {
		ctx.WriteString("usingClause")
		node.UsingClause.Fingerprint(ctx, "UsingClause")
	}
}

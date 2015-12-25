// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node JoinExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("JoinExpr")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx, "Alias")
	}

	ctx.WriteString(strconv.FormatBool(node.IsNatural))
	ctx.WriteString(strconv.Itoa(int(node.Jointype)))

	if node.Larg != nil {
		node.Larg.Fingerprint(ctx, "Larg")
	}

	if node.Quals != nil {
		node.Quals.Fingerprint(ctx, "Quals")
	}

	if node.Rarg != nil {
		node.Rarg.Fingerprint(ctx, "Rarg")
	}

	ctx.WriteString(strconv.Itoa(int(node.Rtindex)))
	node.UsingClause.Fingerprint(ctx, "UsingClause")
}

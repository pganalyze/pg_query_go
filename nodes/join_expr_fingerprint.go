// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node JoinExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("JOINEXPR")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.IsNatural))
	ctx.WriteString(strconv.Itoa(int(node.Jointype)))

	if node.Larg != nil {
		node.Larg.Fingerprint(ctx)
	}

	if node.Quals != nil {
		node.Quals.Fingerprint(ctx)
	}

	if node.Rarg != nil {
		node.Rarg.Fingerprint(ctx)
	}

	for _, subNode := range node.UsingClause {
		subNode.Fingerprint(ctx)
	}
}

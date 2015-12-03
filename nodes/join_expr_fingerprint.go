// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node JoinExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "JoinExpr")
	if node.Alias != nil {
		node.Alias.Fingerprint(ctx)
	}

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

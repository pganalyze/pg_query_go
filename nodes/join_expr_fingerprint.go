// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node JoinExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "JOINEXPR")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.IsNatural))
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Jointype)))

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

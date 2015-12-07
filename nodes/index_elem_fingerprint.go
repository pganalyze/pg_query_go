// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node IndexElem) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("INDEXELEM")

	for _, subNode := range node.Collation {
		subNode.Fingerprint(ctx)
	}

	if node.Expr != nil {
		node.Expr.Fingerprint(ctx)
	}

	if node.Indexcolname != nil {
		ctx.WriteString(*node.Indexcolname)
	}

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}

	ctx.WriteString(strconv.Itoa(int(node.NullsOrdering)))

	for _, subNode := range node.Opclass {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Ordering)))
}

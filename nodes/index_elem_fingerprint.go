// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node IndexElem) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("IndexElem")
	node.Collation.Fingerprint(ctx, "Collation")

	if node.Expr != nil {
		node.Expr.Fingerprint(ctx, "Expr")
	}

	if node.Indexcolname != nil {
		ctx.WriteString(*node.Indexcolname)
	}

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}

	ctx.WriteString(strconv.Itoa(int(node.NullsOrdering)))
	node.Opclass.Fingerprint(ctx, "Opclass")
	ctx.WriteString(strconv.Itoa(int(node.Ordering)))
}

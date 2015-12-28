// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node IndexElem) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("IndexElem")
	if len(node.Collation.Items) > 0 {
		ctx.WriteString("collation")
		node.Collation.Fingerprint(ctx, "Collation")
	}

	if node.Expr != nil {
		ctx.WriteString("expr")
		node.Expr.Fingerprint(ctx, "Expr")
	}

	if node.Indexcolname != nil {
		ctx.WriteString("indexcolname")
		ctx.WriteString(*node.Indexcolname)
	}

	if node.Name != nil {
		ctx.WriteString("name")
		ctx.WriteString(*node.Name)
	}

	if int(node.NullsOrdering) != 0 {
		ctx.WriteString("nulls_ordering")
		ctx.WriteString(strconv.Itoa(int(node.NullsOrdering)))
	}

	if len(node.Opclass.Items) > 0 {
		ctx.WriteString("opclass")
		node.Opclass.Fingerprint(ctx, "Opclass")
	}

	if int(node.Ordering) != 0 {
		ctx.WriteString("ordering")
		ctx.WriteString(strconv.Itoa(int(node.Ordering)))
	}
}

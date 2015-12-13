// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ArrayRef) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ArrayRef")
	ctx.WriteString(strconv.Itoa(int(node.Refarraytype)))

	if node.Refassgnexpr != nil {
		node.Refassgnexpr.Fingerprint(ctx, "Refassgnexpr")
	}

	ctx.WriteString(strconv.Itoa(int(node.Refcollid)))
	ctx.WriteString(strconv.Itoa(int(node.Refelemtype)))

	if node.Refexpr != nil {
		node.Refexpr.Fingerprint(ctx, "Refexpr")
	}

	for _, subNode := range node.Reflowerindexpr {
		subNode.Fingerprint(ctx, "Reflowerindexpr")
	}

	ctx.WriteString(strconv.Itoa(int(node.Reftypmod)))

	for _, subNode := range node.Refupperindexpr {
		subNode.Fingerprint(ctx, "Refupperindexpr")
	}

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

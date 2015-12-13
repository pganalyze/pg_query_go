// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RowExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RowExpr")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx, "Args")
	}

	for _, subNode := range node.Colnames {
		subNode.Fingerprint(ctx, "Colnames")
	}

	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.RowFormat)))
	ctx.WriteString(strconv.Itoa(int(node.RowTypeid)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

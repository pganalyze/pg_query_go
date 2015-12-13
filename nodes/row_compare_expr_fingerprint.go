// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RowCompareExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RowCompareExpr")

	for _, subNode := range node.Inputcollids {
		subNode.Fingerprint(ctx, "Inputcollids")
	}

	for _, subNode := range node.Largs {
		subNode.Fingerprint(ctx, "Largs")
	}

	for _, subNode := range node.Opfamilies {
		subNode.Fingerprint(ctx, "Opfamilies")
	}

	for _, subNode := range node.Opnos {
		subNode.Fingerprint(ctx, "Opnos")
	}

	for _, subNode := range node.Rargs {
		subNode.Fingerprint(ctx, "Rargs")
	}

	ctx.WriteString(strconv.Itoa(int(node.Rctype)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RowCompareExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RowCompareExpr")

	for _, subNode := range node.Inputcollids {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Largs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Opfamilies {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Opnos {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Rargs {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Rctype)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}

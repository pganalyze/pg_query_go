// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node RowCompareExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RowCompareExpr")

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
}

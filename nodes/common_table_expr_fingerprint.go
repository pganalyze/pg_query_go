// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CommonTableExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CommonTableExpr")

	for _, subNode := range node.Aliascolnames {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Ctecolcollations {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Ctecolnames {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Ctecoltypes {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Ctecoltypmods {
		subNode.Fingerprint(ctx)
	}

	if node.Ctequery != nil {
		node.Ctequery.Fingerprint(ctx)
	}
}

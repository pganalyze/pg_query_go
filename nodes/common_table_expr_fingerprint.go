// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CommonTableExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("COMMONTABLEEXPR")

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

	if node.Ctename != nil {
		ctx.WriteString(*node.Ctename)
	}

	if node.Ctequery != nil {
		node.Ctequery.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Cterecursive))
	// Intentionally ignoring node.Location for fingerprinting
}

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CommonTableExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CommonTableExpr")

	for _, subNode := range node.Aliascolnames {
		subNode.Fingerprint(ctx, "Aliascolnames")
	}

	for _, subNode := range node.Ctecolcollations {
		subNode.Fingerprint(ctx, "Ctecolcollations")
	}

	for _, subNode := range node.Ctecolnames {
		subNode.Fingerprint(ctx, "Ctecolnames")
	}

	for _, subNode := range node.Ctecoltypes {
		subNode.Fingerprint(ctx, "Ctecoltypes")
	}

	for _, subNode := range node.Ctecoltypmods {
		subNode.Fingerprint(ctx, "Ctecoltypmods")
	}

	if node.Ctename != nil {
		ctx.WriteString(*node.Ctename)
	}

	if node.Ctequery != nil {
		node.Ctequery.Fingerprint(ctx, "Ctequery")
	}

	ctx.WriteString(strconv.FormatBool(node.Cterecursive))
	ctx.WriteString(strconv.Itoa(int(node.Cterefcount)))
	// Intentionally ignoring node.Location for fingerprinting
}

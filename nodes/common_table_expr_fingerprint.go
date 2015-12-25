// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CommonTableExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CommonTableExpr")
	node.Aliascolnames.Fingerprint(ctx, "Aliascolnames")
	node.Ctecolcollations.Fingerprint(ctx, "Ctecolcollations")
	node.Ctecolnames.Fingerprint(ctx, "Ctecolnames")
	node.Ctecoltypes.Fingerprint(ctx, "Ctecoltypes")
	node.Ctecoltypmods.Fingerprint(ctx, "Ctecoltypmods")

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

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CommonTableExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CommonTableExpr")
	if len(node.Aliascolnames.Items) > 0 {
		ctx.WriteString("aliascolnames")
		node.Aliascolnames.Fingerprint(ctx, "Aliascolnames")
	}

	if len(node.Ctecolcollations.Items) > 0 {
		ctx.WriteString("ctecolcollations")
		node.Ctecolcollations.Fingerprint(ctx, "Ctecolcollations")
	}

	if len(node.Ctecolnames.Items) > 0 {
		ctx.WriteString("ctecolnames")
		node.Ctecolnames.Fingerprint(ctx, "Ctecolnames")
	}

	if len(node.Ctecoltypes.Items) > 0 {
		ctx.WriteString("ctecoltypes")
		node.Ctecoltypes.Fingerprint(ctx, "Ctecoltypes")
	}

	if len(node.Ctecoltypmods.Items) > 0 {
		ctx.WriteString("ctecoltypmods")
		node.Ctecoltypmods.Fingerprint(ctx, "Ctecoltypmods")
	}

	if node.Ctename != nil {
		ctx.WriteString("ctename")
		ctx.WriteString(*node.Ctename)
	}

	if node.Ctequery != nil {
		ctx.WriteString("ctequery")
		node.Ctequery.Fingerprint(ctx, "Ctequery")
	}

	if node.Cterecursive {
		ctx.WriteString("cterecursive")
		ctx.WriteString(strconv.FormatBool(node.Cterecursive))
	}

	if node.Cterefcount != 0 {
		ctx.WriteString("cterefcount")
		ctx.WriteString(strconv.Itoa(int(node.Cterefcount)))
	}

	// Intentionally ignoring node.Location for fingerprinting
}

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeTblFunction) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RangeTblFunction")
	if len(node.Funccolcollations.Items) > 0 {
		ctx.WriteString("funccolcollations")
		node.Funccolcollations.Fingerprint(ctx, "Funccolcollations")
	}

	if node.Funccolcount != 0 {
		ctx.WriteString("funccolcount")
		ctx.WriteString(strconv.Itoa(int(node.Funccolcount)))
	}

	if len(node.Funccolnames.Items) > 0 {
		ctx.WriteString("funccolnames")
		node.Funccolnames.Fingerprint(ctx, "Funccolnames")
	}

	if len(node.Funccoltypes.Items) > 0 {
		ctx.WriteString("funccoltypes")
		node.Funccoltypes.Fingerprint(ctx, "Funccoltypes")
	}

	if len(node.Funccoltypmods.Items) > 0 {
		ctx.WriteString("funccoltypmods")
		node.Funccoltypmods.Fingerprint(ctx, "Funccoltypmods")
	}

	if node.Funcexpr != nil {
		ctx.WriteString("funcexpr")
		node.Funcexpr.Fingerprint(ctx, "Funcexpr")
	}

	ctx.WriteString("funcparams")
	for _, val := range node.Funcparams {
		ctx.WriteString(strconv.Itoa(int(val)))
	}
}

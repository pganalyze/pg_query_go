// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeTblFunction) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RangeTblFunction")
	node.Funccolcollations.Fingerprint(ctx, "Funccolcollations")
	ctx.WriteString(strconv.Itoa(int(node.Funccolcount)))
	node.Funccolnames.Fingerprint(ctx, "Funccolnames")
	node.Funccoltypes.Fingerprint(ctx, "Funccoltypes")
	node.Funccoltypmods.Fingerprint(ctx, "Funccoltypmods")

	if node.Funcexpr != nil {
		node.Funcexpr.Fingerprint(ctx, "Funcexpr")
	}

	for _, val := range node.Funcparams {
		ctx.WriteString(strconv.Itoa(int(val)))
	}
}

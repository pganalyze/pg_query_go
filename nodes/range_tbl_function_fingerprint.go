// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeTblFunction) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RangeTblFunction")

	for _, subNode := range node.Funccolcollations {
		subNode.Fingerprint(ctx, "Funccolcollations")
	}

	ctx.WriteString(strconv.Itoa(int(node.Funccolcount)))

	for _, subNode := range node.Funccolnames {
		subNode.Fingerprint(ctx, "Funccolnames")
	}

	for _, subNode := range node.Funccoltypes {
		subNode.Fingerprint(ctx, "Funccoltypes")
	}

	for _, subNode := range node.Funccoltypmods {
		subNode.Fingerprint(ctx, "Funccoltypmods")
	}

	if node.Funcexpr != nil {
		node.Funcexpr.Fingerprint(ctx, "Funcexpr")
	}

	for _, val := range node.Funcparams {
		ctx.WriteString(strconv.Itoa(int(val)))
	}
}

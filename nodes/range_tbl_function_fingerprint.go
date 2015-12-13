// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeTblFunction) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RangeTblFunction")

	for _, subNode := range node.Funccolcollations {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Funccolcount)))

	for _, subNode := range node.Funccolnames {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Funccoltypes {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Funccoltypmods {
		subNode.Fingerprint(ctx)
	}

	if node.Funcexpr != nil {
		node.Funcexpr.Fingerprint(ctx)
	}

	for _, val := range node.Funcparams {
		ctx.WriteString(strconv.Itoa(int(val)))
	}
}

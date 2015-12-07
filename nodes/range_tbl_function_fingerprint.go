// Auto-generated - DO NOT EDIT

package pg_query

func (node RangeTblFunction) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RANGETBLFUNCTION")

	for _, subNode := range node.Funccolcollations {
		subNode.Fingerprint(ctx)
	}

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
}

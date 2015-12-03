// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node FuncCall) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FuncCall")
	if node.AggFilter != nil {
		node.AggFilter.Fingerprint(ctx)
	}

	for _, subNode := range node.AggOrder {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Funcname {
		subNode.Fingerprint(ctx)
	}

	if node.Over != nil {
		node.Over.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node SubPlan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SubPlan")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.ParParam {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.ParamIds {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.SetParam {
		subNode.Fingerprint(ctx)
	}

	if node.Testexpr != nil {
		node.Testexpr.Fingerprint(ctx)
	}
}

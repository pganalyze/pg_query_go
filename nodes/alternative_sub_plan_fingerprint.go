// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlternativeSubPlan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlternativeSubPlan")

	for _, subNode := range node.Subplans {
		subNode.Fingerprint(ctx)
	}
}

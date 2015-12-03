// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node PlannerParamItem) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PlannerParamItem")
	if node.Item != nil {
		node.Item.Fingerprint(ctx)
	}
}

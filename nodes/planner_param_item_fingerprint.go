// Auto-generated - DO NOT EDIT

package pg_query

func (node PlannerParamItem) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("PLANNERPARAMITEM")

	if node.Item != nil {
		node.Item.Fingerprint(ctx)
	}
}

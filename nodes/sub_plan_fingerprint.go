// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SubPlan) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("SUBPLAN")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.ParParam {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.ParamIds {
		subNode.Fingerprint(ctx)
	}

	if node.PlanName != nil {
		ctx.WriteString(*node.PlanName)
	}

	for _, subNode := range node.SetParam {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.SubLinkType)))

	if node.Testexpr != nil {
		node.Testexpr.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.UnknownEqFalse))
	ctx.WriteString(strconv.FormatBool(node.UseHashTable))
}

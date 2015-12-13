// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SubPlan) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SubPlan")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx, "Args")
	}

	ctx.WriteString(strconv.Itoa(int(node.FirstColCollation)))
	ctx.WriteString(strconv.Itoa(int(node.FirstColType)))
	ctx.WriteString(strconv.Itoa(int(node.FirstColTypmod)))

	for _, subNode := range node.ParParam {
		subNode.Fingerprint(ctx, "ParParam")
	}

	for _, subNode := range node.ParamIds {
		subNode.Fingerprint(ctx, "ParamIds")
	}

	ctx.WriteString(strconv.FormatFloat(float64(node.PerCallCost), 'E', -1, 64))
	ctx.WriteString(strconv.Itoa(int(node.PlanId)))

	if node.PlanName != nil {
		ctx.WriteString(*node.PlanName)
	}

	for _, subNode := range node.SetParam {
		subNode.Fingerprint(ctx, "SetParam")
	}

	ctx.WriteString(strconv.FormatFloat(float64(node.StartupCost), 'E', -1, 64))
	ctx.WriteString(strconv.Itoa(int(node.SubLinkType)))

	if node.Testexpr != nil {
		node.Testexpr.Fingerprint(ctx, "Testexpr")
	}

	ctx.WriteString(strconv.FormatBool(node.UnknownEqFalse))
	ctx.WriteString(strconv.FormatBool(node.UseHashTable))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

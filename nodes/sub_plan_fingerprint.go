// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SubPlan) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SubPlan")
	node.Args.Fingerprint(ctx, "Args")
	ctx.WriteString(strconv.Itoa(int(node.FirstColCollation)))
	ctx.WriteString(strconv.Itoa(int(node.FirstColType)))
	ctx.WriteString(strconv.Itoa(int(node.FirstColTypmod)))
	node.ParParam.Fingerprint(ctx, "ParParam")
	node.ParamIds.Fingerprint(ctx, "ParamIds")
	ctx.WriteString(strconv.FormatFloat(float64(node.PerCallCost), 'E', -1, 64))
	ctx.WriteString(strconv.Itoa(int(node.PlanId)))

	if node.PlanName != nil {
		ctx.WriteString(*node.PlanName)
	}

	node.SetParam.Fingerprint(ctx, "SetParam")
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

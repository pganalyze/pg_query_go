// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SubPlan) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SubPlan")
	if len(node.Args.Items) > 0 {
		ctx.WriteString("args")
		node.Args.Fingerprint(ctx, "Args")
	}

	if node.FirstColCollation != 0 {
		ctx.WriteString("firstColCollation")
		ctx.WriteString(strconv.Itoa(int(node.FirstColCollation)))
	}

	if node.FirstColType != 0 {
		ctx.WriteString("firstColType")
		ctx.WriteString(strconv.Itoa(int(node.FirstColType)))
	}

	if node.FirstColTypmod != 0 {
		ctx.WriteString("firstColTypmod")
		ctx.WriteString(strconv.Itoa(int(node.FirstColTypmod)))
	}

	if len(node.ParParam.Items) > 0 {
		ctx.WriteString("parParam")
		node.ParParam.Fingerprint(ctx, "ParParam")
	}

	if len(node.ParamIds.Items) > 0 {
		ctx.WriteString("paramIds")
		node.ParamIds.Fingerprint(ctx, "ParamIds")
	}

	ctx.WriteString("per_call_cost")
	ctx.WriteString(strconv.FormatFloat(float64(node.PerCallCost), 'E', -1, 64))

	if node.PlanId != 0 {
		ctx.WriteString("plan_id")
		ctx.WriteString(strconv.Itoa(int(node.PlanId)))
	}

	if node.PlanName != nil {
		ctx.WriteString("plan_name")
		ctx.WriteString(*node.PlanName)
	}

	if len(node.SetParam.Items) > 0 {
		ctx.WriteString("setParam")
		node.SetParam.Fingerprint(ctx, "SetParam")
	}

	ctx.WriteString("startup_cost")
	ctx.WriteString(strconv.FormatFloat(float64(node.StartupCost), 'E', -1, 64))

	if int(node.SubLinkType) != 0 {
		ctx.WriteString("subLinkType")
		ctx.WriteString(strconv.Itoa(int(node.SubLinkType)))
	}

	if node.Testexpr != nil {
		ctx.WriteString("testexpr")
		node.Testexpr.Fingerprint(ctx, "Testexpr")
	}

	if node.UnknownEqFalse {
		ctx.WriteString("unknownEqFalse")
		ctx.WriteString(strconv.FormatBool(node.UnknownEqFalse))
	}

	if node.UseHashTable {
		ctx.WriteString("useHashTable")
		ctx.WriteString(strconv.FormatBool(node.UseHashTable))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

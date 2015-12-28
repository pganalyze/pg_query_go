// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node FuncCall) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FuncCall")

	if node.AggDistinct {
		ctx.WriteString("agg_distinct")
		ctx.WriteString(strconv.FormatBool(node.AggDistinct))
	}

	if node.AggFilter != nil {
		ctx.WriteString("agg_filter")
		node.AggFilter.Fingerprint(ctx, "AggFilter")
	}

	if len(node.AggOrder.Items) > 0 {
		ctx.WriteString("agg_order")
		node.AggOrder.Fingerprint(ctx, "AggOrder")
	}

	if node.AggStar {
		ctx.WriteString("agg_star")
		ctx.WriteString(strconv.FormatBool(node.AggStar))
	}

	if node.AggWithinGroup {
		ctx.WriteString("agg_within_group")
		ctx.WriteString(strconv.FormatBool(node.AggWithinGroup))
	}

	if len(node.Args.Items) > 0 {
		ctx.WriteString("args")
		node.Args.Fingerprint(ctx, "Args")
	}

	if node.FuncVariadic {
		ctx.WriteString("func_variadic")
		ctx.WriteString(strconv.FormatBool(node.FuncVariadic))
	}

	if len(node.Funcname.Items) > 0 {
		ctx.WriteString("funcname")
		node.Funcname.Fingerprint(ctx, "Funcname")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Over != nil {
		ctx.WriteString("over")
		node.Over.Fingerprint(ctx, "Over")
	}
}

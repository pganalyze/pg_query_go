// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node FuncCall) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FuncCall")
	ctx.WriteString(strconv.FormatBool(node.AggDistinct))

	if node.AggFilter != nil {
		node.AggFilter.Fingerprint(ctx, "AggFilter")
	}

	for _, subNode := range node.AggOrder {
		subNode.Fingerprint(ctx, "AggOrder")
	}

	ctx.WriteString(strconv.FormatBool(node.AggStar))
	ctx.WriteString(strconv.FormatBool(node.AggWithinGroup))

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx, "Args")
	}

	ctx.WriteString(strconv.FormatBool(node.FuncVariadic))

	for _, subNode := range node.Funcname {
		subNode.Fingerprint(ctx, "Funcname")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Over != nil {
		node.Over.Fingerprint(ctx, "Over")
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node SubPlan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SUBPLAN")

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
		io.WriteString(ctx.hash, *node.PlanName)
	}

	for _, subNode := range node.SetParam {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.SubLinkType)))

	if node.Testexpr != nil {
		node.Testexpr.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.UnknownEqFalse))
	io.WriteString(ctx.hash, strconv.FormatBool(node.UseHashTable))
}

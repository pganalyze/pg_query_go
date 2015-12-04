// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node PlannerGlobal) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PLANNERGLOBAL")

	for _, subNode := range node.Finalrowmarks {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Finalrtable {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.InvalItems {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.RelationOids {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.ResultRelations {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Subplans {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Subroots {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.TransientPlan))
}

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node PlannerGlobal) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("PLANNERGLOBAL")

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

	ctx.WriteString(strconv.FormatBool(node.TransientPlan))
}

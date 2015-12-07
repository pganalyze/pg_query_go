// Auto-generated - DO NOT EDIT

package pg_query

func (node Plan) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("PLAN")

	for _, subNode := range node.InitPlan {
		subNode.Fingerprint(ctx)
	}

	if node.Lefttree != nil {
		node.Lefttree.Fingerprint(ctx)
	}

	for _, subNode := range node.Qual {
		subNode.Fingerprint(ctx)
	}

	if node.Righttree != nil {
		node.Righttree.Fingerprint(ctx)
	}

	for _, subNode := range node.Targetlist {
		subNode.Fingerprint(ctx)
	}
}

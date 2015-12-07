// Auto-generated - DO NOT EDIT

package pg_query

func (node Path) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("PATH")

	if node.ParamInfo != nil {
		node.ParamInfo.Fingerprint(ctx)
	}

	if node.Parent != nil {
		node.Parent.Fingerprint(ctx)
	}

	for _, subNode := range node.Pathkeys {
		subNode.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateOpClassItem) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CREATEOPCLASSITEM")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.ClassArgs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Name {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.OrderFamily {
		subNode.Fingerprint(ctx)
	}

	if node.Storedtype != nil {
		node.Storedtype.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

func (node FieldStore) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("FIELDSTORE")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	for _, subNode := range node.Fieldnums {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Newvals {
		subNode.Fingerprint(ctx)
	}
}

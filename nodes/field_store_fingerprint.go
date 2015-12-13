// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node FieldStore) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("FieldStore")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	for _, subNode := range node.Fieldnums {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Newvals {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Resulttype)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}

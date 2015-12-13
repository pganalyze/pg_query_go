// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node FieldStore) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FieldStore")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx, "Arg")
	}

	for _, subNode := range node.Fieldnums {
		subNode.Fingerprint(ctx, "Fieldnums")
	}

	for _, subNode := range node.Newvals {
		subNode.Fingerprint(ctx, "Newvals")
	}

	ctx.WriteString(strconv.Itoa(int(node.Resulttype)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

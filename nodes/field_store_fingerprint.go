// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node FieldStore) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FIELDSTORE")

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

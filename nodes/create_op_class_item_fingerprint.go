// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateOpClassItem) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateOpClassItem")

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

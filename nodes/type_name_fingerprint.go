// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node TypeName) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "TypeName")

	for _, subNode := range node.ArrayBounds {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Names {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Typmods {
		subNode.Fingerprint(ctx)
	}
}

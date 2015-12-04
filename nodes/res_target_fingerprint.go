// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ResTarget) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RESTARGET")

	for _, subNode := range node.Indirection {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}

	if node.Val != nil {
		node.Val.Fingerprint(ctx)
	}
}

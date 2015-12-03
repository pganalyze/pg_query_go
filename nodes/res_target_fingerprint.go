// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ResTarget) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ResTarget")

	for _, subNode := range node.Indirection {
		subNode.Fingerprint(ctx)
	}

	if node.Val != nil {
		node.Val.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node A_Indirection) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "A_Indirection")
	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	for _, subNode := range node.Indirection {
		subNode.Fingerprint(ctx)
	}
}

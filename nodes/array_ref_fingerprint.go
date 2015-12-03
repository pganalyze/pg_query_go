// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ArrayRef) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ArrayRef")
	if node.Refassgnexpr != nil {
		node.Refassgnexpr.Fingerprint(ctx)
	}

	if node.Refexpr != nil {
		node.Refexpr.Fingerprint(ctx)
	}

	for _, subNode := range node.Reflowerindexpr {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Refupperindexpr {
		subNode.Fingerprint(ctx)
	}
}

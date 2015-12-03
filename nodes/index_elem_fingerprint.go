// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node IndexElem) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "IndexElem")

	for _, subNode := range node.Collation {
		subNode.Fingerprint(ctx)
	}

	if node.Expr != nil {
		node.Expr.Fingerprint(ctx)
	}

	for _, subNode := range node.Opclass {
		subNode.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CollateClause) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CollateClause")
	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	for _, subNode := range node.Collname {
		subNode.Fingerprint(ctx)
	}
}

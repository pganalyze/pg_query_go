// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node IndexOptInfo) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "IndexOptInfo")

	for _, subNode := range node.Indexprs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indextlist {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indpred {
		subNode.Fingerprint(ctx)
	}

	if node.Rel != nil {
		node.Rel.Fingerprint(ctx)
	}
}

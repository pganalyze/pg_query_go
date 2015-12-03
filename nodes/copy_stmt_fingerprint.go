// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CopyStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CopyStmt")

	for _, subNode := range node.Attlist {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}
}

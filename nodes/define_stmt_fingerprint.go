// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node DefineStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DefineStmt")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Definition {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Defnames {
		subNode.Fingerprint(ctx)
	}
}

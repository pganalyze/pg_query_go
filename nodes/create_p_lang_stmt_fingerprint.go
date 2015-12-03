// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreatePLangStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreatePLangStmt")

	for _, subNode := range node.Plhandler {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Plinline {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Plvalidator {
		subNode.Fingerprint(ctx)
	}
}

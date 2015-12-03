// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node SecLabelStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SecLabelStmt")

	for _, subNode := range node.Objargs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Objname {
		subNode.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterExtensionContentsStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterExtensionContentsStmt")

	for _, subNode := range node.Objargs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Objname {
		subNode.Fingerprint(ctx)
	}
}

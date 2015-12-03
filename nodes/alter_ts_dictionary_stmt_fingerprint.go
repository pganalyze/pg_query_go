// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterTSDictionaryStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterTSDictionaryStmt")

	for _, subNode := range node.Dictname {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterUserMappingStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTERUSERMAPPINGSTMT")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Servername != nil {
		io.WriteString(ctx.hash, *node.Servername)
	}

	if node.Username != nil {
		io.WriteString(ctx.hash, *node.Username)
	}
}

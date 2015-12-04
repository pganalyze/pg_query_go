// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateTableSpaceStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATETABLESPACESTMT")
	// Intentionally ignoring node.Location for fingerprinting

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Owner != nil {
		io.WriteString(ctx.hash, *node.Owner)
	}

	if node.Tablespacename != nil {
		io.WriteString(ctx.hash, *node.Tablespacename)
	}
}

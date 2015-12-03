// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterDefaultPrivilegesStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterDefaultPrivilegesStmt")
	if node.Action != nil {
		node.Action.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}

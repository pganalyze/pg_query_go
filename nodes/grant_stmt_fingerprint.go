// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node GrantStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "GrantStmt")

	for _, subNode := range node.Grantees {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Objects {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Privileges {
		subNode.Fingerprint(ctx)
	}
}

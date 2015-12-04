// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ReplicaIdentityStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "REPLICAIDENTITYSTMT")

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}
}

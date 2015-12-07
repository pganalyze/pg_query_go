// Auto-generated - DO NOT EDIT

package pg_query

func (node ReplicaIdentityStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("REPLICAIDENTITYSTMT")

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}
}

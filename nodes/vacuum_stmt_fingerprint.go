// Auto-generated - DO NOT EDIT

package pg_query

func (node VacuumStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("VACUUM")

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	for _, subNode := range node.VaCols {
		subNode.Fingerprint(ctx)
	}
}

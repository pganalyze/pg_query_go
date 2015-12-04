// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node VacuumStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "VACUUM")

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	for _, subNode := range node.VaCols {
		subNode.Fingerprint(ctx)
	}
}

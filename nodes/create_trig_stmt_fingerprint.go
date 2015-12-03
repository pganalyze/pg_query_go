// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateTrigStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateTrigStmt")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Columns {
		subNode.Fingerprint(ctx)
	}

	if node.Constrrel != nil {
		node.Constrrel.Fingerprint(ctx)
	}

	for _, subNode := range node.Funcname {
		subNode.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	if node.WhenClause != nil {
		node.WhenClause.Fingerprint(ctx)
	}
}

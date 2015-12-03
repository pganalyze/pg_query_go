// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateEventTrigStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateEventTrigStmt")

	for _, subNode := range node.Funcname {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Whenclause {
		subNode.Fingerprint(ctx)
	}
}

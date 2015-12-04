// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateEventTrigStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATEEVENTTRIGSTMT")

	if node.Eventname != nil {
		io.WriteString(ctx.hash, *node.Eventname)
	}

	for _, subNode := range node.Funcname {
		subNode.Fingerprint(ctx)
	}

	if node.Trigname != nil {
		io.WriteString(ctx.hash, *node.Trigname)
	}

	for _, subNode := range node.Whenclause {
		subNode.Fingerprint(ctx)
	}
}

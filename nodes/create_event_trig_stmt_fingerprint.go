// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateEventTrigStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CREATEEVENTTRIGSTMT")

	if node.Eventname != nil {
		ctx.WriteString(*node.Eventname)
	}

	for _, subNode := range node.Funcname {
		subNode.Fingerprint(ctx)
	}

	if node.Trigname != nil {
		ctx.WriteString(*node.Trigname)
	}

	for _, subNode := range node.Whenclause {
		subNode.Fingerprint(ctx)
	}
}

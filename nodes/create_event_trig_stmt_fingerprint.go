// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateEventTrigStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateEventTrigStmt")

	if node.Eventname != nil {
		ctx.WriteString(*node.Eventname)
	}

	for _, subNode := range node.Funcname {
		subNode.Fingerprint(ctx, "Funcname")
	}

	if node.Trigname != nil {
		ctx.WriteString(*node.Trigname)
	}

	for _, subNode := range node.Whenclause {
		subNode.Fingerprint(ctx, "Whenclause")
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateEventTrigStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateEventTrigStmt")

	if node.Eventname != nil {
		ctx.WriteString(*node.Eventname)
	}

	node.Funcname.Fingerprint(ctx, "Funcname")

	if node.Trigname != nil {
		ctx.WriteString(*node.Trigname)
	}

	node.Whenclause.Fingerprint(ctx, "Whenclause")
}

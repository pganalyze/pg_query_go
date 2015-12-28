// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateEventTrigStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateEventTrigStmt")

	if node.Eventname != nil {
		ctx.WriteString("eventname")
		ctx.WriteString(*node.Eventname)
	}

	if len(node.Funcname.Items) > 0 {
		ctx.WriteString("funcname")
		node.Funcname.Fingerprint(ctx, "Funcname")
	}

	if node.Trigname != nil {
		ctx.WriteString("trigname")
		ctx.WriteString(*node.Trigname)
	}

	if len(node.Whenclause.Items) > 0 {
		ctx.WriteString("whenclause")
		node.Whenclause.Fingerprint(ctx, "Whenclause")
	}
}

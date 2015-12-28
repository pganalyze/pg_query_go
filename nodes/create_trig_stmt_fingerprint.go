// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateTrigStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateTrigStmt")
	if len(node.Args.Items) > 0 {
		ctx.WriteString("args")
		node.Args.Fingerprint(ctx, "Args")
	}

	if len(node.Columns.Items) > 0 {
		ctx.WriteString("columns")
		node.Columns.Fingerprint(ctx, "Columns")
	}

	if node.Constrrel != nil {
		ctx.WriteString("constrrel")
		node.Constrrel.Fingerprint(ctx, "Constrrel")
	}

	if node.Deferrable {
		ctx.WriteString("deferrable")
		ctx.WriteString(strconv.FormatBool(node.Deferrable))
	}

	if node.Events != 0 {
		ctx.WriteString("events")
		ctx.WriteString(strconv.Itoa(int(node.Events)))
	}

	if len(node.Funcname.Items) > 0 {
		ctx.WriteString("funcname")
		node.Funcname.Fingerprint(ctx, "Funcname")
	}

	if node.Initdeferred {
		ctx.WriteString("initdeferred")
		ctx.WriteString(strconv.FormatBool(node.Initdeferred))
	}

	if node.Isconstraint {
		ctx.WriteString("isconstraint")
		ctx.WriteString(strconv.FormatBool(node.Isconstraint))
	}

	if node.Relation != nil {
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}

	if node.Row {
		ctx.WriteString("row")
		ctx.WriteString(strconv.FormatBool(node.Row))
	}

	if node.Timing != 0 {
		ctx.WriteString("timing")
		ctx.WriteString(strconv.Itoa(int(node.Timing)))
	}

	if node.Trigname != nil {
		ctx.WriteString("trigname")
		ctx.WriteString(*node.Trigname)
	}

	if node.WhenClause != nil {
		ctx.WriteString("whenClause")
		node.WhenClause.Fingerprint(ctx, "WhenClause")
	}
}

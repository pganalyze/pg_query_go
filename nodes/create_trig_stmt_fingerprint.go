// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateTrigStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateTrigStmt")
	node.Args.Fingerprint(ctx, "Args")
	node.Columns.Fingerprint(ctx, "Columns")

	if node.Constrrel != nil {
		node.Constrrel.Fingerprint(ctx, "Constrrel")
	}

	ctx.WriteString(strconv.FormatBool(node.Deferrable))
	ctx.WriteString(strconv.Itoa(int(node.Events)))
	node.Funcname.Fingerprint(ctx, "Funcname")
	ctx.WriteString(strconv.FormatBool(node.Initdeferred))
	ctx.WriteString(strconv.FormatBool(node.Isconstraint))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}

	ctx.WriteString(strconv.FormatBool(node.Row))
	ctx.WriteString(strconv.Itoa(int(node.Timing)))

	if node.Trigname != nil {
		ctx.WriteString(*node.Trigname)
	}

	if node.WhenClause != nil {
		node.WhenClause.Fingerprint(ctx, "WhenClause")
	}
}

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateTrigStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CreateTrigStmt")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Columns {
		subNode.Fingerprint(ctx)
	}

	if node.Constrrel != nil {
		node.Constrrel.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Deferrable))
	ctx.WriteString(strconv.Itoa(int(node.Events)))

	for _, subNode := range node.Funcname {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Initdeferred))
	ctx.WriteString(strconv.FormatBool(node.Isconstraint))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Row))
	ctx.WriteString(strconv.Itoa(int(node.Timing)))

	if node.Trigname != nil {
		ctx.WriteString(*node.Trigname)
	}

	if node.WhenClause != nil {
		node.WhenClause.Fingerprint(ctx)
	}
}

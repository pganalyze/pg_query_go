// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SetOperationStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SetOperationStmt")

	if node.All {
		ctx.WriteString("all")
		ctx.WriteString(strconv.FormatBool(node.All))
	}

	if len(node.ColCollations.Items) > 0 {
		ctx.WriteString("colCollations")
		node.ColCollations.Fingerprint(ctx, "ColCollations")
	}

	if len(node.ColTypes.Items) > 0 {
		ctx.WriteString("colTypes")
		node.ColTypes.Fingerprint(ctx, "ColTypes")
	}

	if len(node.ColTypmods.Items) > 0 {
		ctx.WriteString("colTypmods")
		node.ColTypmods.Fingerprint(ctx, "ColTypmods")
	}

	if len(node.GroupClauses.Items) > 0 {
		ctx.WriteString("groupClauses")
		node.GroupClauses.Fingerprint(ctx, "GroupClauses")
	}

	if node.Larg != nil {
		ctx.WriteString("larg")
		node.Larg.Fingerprint(ctx, "Larg")
	}

	if int(node.Op) != 0 {
		ctx.WriteString("op")
		ctx.WriteString(strconv.Itoa(int(node.Op)))
	}

	if node.Rarg != nil {
		ctx.WriteString("rarg")
		node.Rarg.Fingerprint(ctx, "Rarg")
	}
}

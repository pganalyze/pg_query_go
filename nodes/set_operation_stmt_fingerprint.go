// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SetOperationStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SetOperationStmt")
	ctx.WriteString(strconv.FormatBool(node.All))
	node.ColCollations.Fingerprint(ctx, "ColCollations")
	node.ColTypes.Fingerprint(ctx, "ColTypes")
	node.ColTypmods.Fingerprint(ctx, "ColTypmods")
	node.GroupClauses.Fingerprint(ctx, "GroupClauses")

	if node.Larg != nil {
		node.Larg.Fingerprint(ctx, "Larg")
	}

	ctx.WriteString(strconv.Itoa(int(node.Op)))

	if node.Rarg != nil {
		node.Rarg.Fingerprint(ctx, "Rarg")
	}
}

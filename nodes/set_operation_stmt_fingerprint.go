// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SetOperationStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SetOperationStmt")
	ctx.WriteString(strconv.FormatBool(node.All))

	for _, subNode := range node.ColCollations {
		subNode.Fingerprint(ctx, "ColCollations")
	}

	for _, subNode := range node.ColTypes {
		subNode.Fingerprint(ctx, "ColTypes")
	}

	for _, subNode := range node.ColTypmods {
		subNode.Fingerprint(ctx, "ColTypmods")
	}

	for _, subNode := range node.GroupClauses {
		subNode.Fingerprint(ctx, "GroupClauses")
	}

	if node.Larg != nil {
		node.Larg.Fingerprint(ctx, "Larg")
	}

	ctx.WriteString(strconv.Itoa(int(node.Op)))

	if node.Rarg != nil {
		node.Rarg.Fingerprint(ctx, "Rarg")
	}
}

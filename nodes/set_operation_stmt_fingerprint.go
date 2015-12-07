// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SetOperationStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("SETOPERATIONSTMT")
	ctx.WriteString(strconv.FormatBool(node.All))

	for _, subNode := range node.ColCollations {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.ColTypes {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.ColTypmods {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.GroupClauses {
		subNode.Fingerprint(ctx)
	}

	if node.Larg != nil {
		node.Larg.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Op)))

	if node.Rarg != nil {
		node.Rarg.Fingerprint(ctx)
	}
}

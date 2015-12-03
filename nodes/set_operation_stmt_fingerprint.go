// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node SetOperationStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SetOperationStmt")

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

	if node.Rarg != nil {
		node.Rarg.Fingerprint(ctx)
	}
}

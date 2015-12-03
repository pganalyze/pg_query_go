// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterObjectSchemaStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterObjectSchemaStmt")

	for _, subNode := range node.Objarg {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Object {
		subNode.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}
}

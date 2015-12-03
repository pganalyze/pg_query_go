// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateStmt")

	for _, subNode := range node.Constraints {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.InhRelations {
		subNode.Fingerprint(ctx)
	}

	if node.OfTypename != nil {
		node.OfTypename.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	for _, subNode := range node.TableElts {
		subNode.Fingerprint(ctx)
	}
}

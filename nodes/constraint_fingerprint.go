// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node Constraint) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "Constraint")

	for _, subNode := range node.Exclusions {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.FkAttrs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Keys {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.OldConpfeqop {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.PkAttrs {
		subNode.Fingerprint(ctx)
	}

	if node.Pktable != nil {
		node.Pktable.Fingerprint(ctx)
	}

	if node.RawExpr != nil {
		node.RawExpr.Fingerprint(ctx)
	}

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx)
	}
}

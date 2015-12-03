// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ColumnDef) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ColumnDef")
	if node.CollClause != nil {
		node.CollClause.Fingerprint(ctx)
	}

	for _, subNode := range node.Constraints {
		subNode.Fingerprint(ctx)
	}

	if node.CookedDefault != nil {
		node.CookedDefault.Fingerprint(ctx)
	}

	for _, subNode := range node.Fdwoptions {
		subNode.Fingerprint(ctx)
	}

	if node.RawDefault != nil {
		node.RawDefault.Fingerprint(ctx)
	}

	if node.TypeName != nil {
		node.TypeName.Fingerprint(ctx)
	}
}

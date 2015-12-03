// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node EquivalenceMember) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "EquivalenceMember")
	if node.EmExpr != nil {
		node.EmExpr.Fingerprint(ctx)
	}
}

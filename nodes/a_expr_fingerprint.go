// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node A_Expr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("A_Expr")
	ctx.WriteString(strconv.Itoa(int(node.Kind)))

	if node.Lexpr != nil {
		node.Lexpr.Fingerprint(ctx, "Lexpr")
	}

	// Intentionally ignoring node.Location for fingerprinting

	node.Name.Fingerprint(ctx, "Name")

	if node.Rexpr != nil {
		node.Rexpr.Fingerprint(ctx, "Rexpr")
	}
}

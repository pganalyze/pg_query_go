// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node A_Expr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("A_Expr")

	if int(node.Kind) != 0 {
		ctx.WriteString("kind")
		ctx.WriteString(strconv.Itoa(int(node.Kind)))
	}

	if node.Lexpr != nil {
		ctx.WriteString("lexpr")
		node.Lexpr.Fingerprint(ctx, "Lexpr")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if len(node.Name.Items) > 0 {
		ctx.WriteString("name")
		node.Name.Fingerprint(ctx, "Name")
	}

	if node.Rexpr != nil {
		ctx.WriteString("rexpr")
		node.Rexpr.Fingerprint(ctx, "Rexpr")
	}
}

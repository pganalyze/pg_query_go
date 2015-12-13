// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node A_Expr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("A_Expr")
	ctx.WriteString(strconv.Itoa(int(node.Kind)))

	if node.Lexpr != nil {
		node.Lexpr.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	for _, subNode := range node.Name {
		subNode.Fingerprint(ctx)
	}

	if node.Rexpr != nil {
		node.Rexpr.Fingerprint(ctx)
	}
}

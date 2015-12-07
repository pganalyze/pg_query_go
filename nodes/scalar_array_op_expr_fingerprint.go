// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ScalarArrayOpExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("SCALARARRAYOPEXPR")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.FormatBool(node.UseOr))
}

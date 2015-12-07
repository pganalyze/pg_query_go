// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ArrayExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ARRAY")

	for _, subNode := range node.Elements {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.FormatBool(node.Multidims))
}

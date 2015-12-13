// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node BoolExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("BoolExpr")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Boolop)))
	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}

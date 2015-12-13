// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CoalesceExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CoalesceExpr")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Coalescecollid)))
	ctx.WriteString(strconv.Itoa(int(node.Coalescetype)))
	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}

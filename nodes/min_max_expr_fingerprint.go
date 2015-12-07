// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node MinMaxExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("MINMAX")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Op)))
}

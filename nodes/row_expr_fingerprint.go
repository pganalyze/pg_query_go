// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RowExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ROW")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Colnames {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.RowFormat)))
}

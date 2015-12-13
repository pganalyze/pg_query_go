// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node WithClause) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("WithClause")

	for _, subNode := range node.Ctes {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.FormatBool(node.Recursive))
}

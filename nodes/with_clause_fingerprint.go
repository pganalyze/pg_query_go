// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node WithClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("WithClause")
	node.Ctes.Fingerprint(ctx, "Ctes")
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.FormatBool(node.Recursive))
}

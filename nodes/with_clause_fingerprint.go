// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node WithClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("WithClause")
	if len(node.Ctes.Items) > 0 {
		ctx.WriteString("ctes")
		node.Ctes.Fingerprint(ctx, "Ctes")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Recursive {
		ctx.WriteString("recursive")
		ctx.WriteString(strconv.FormatBool(node.Recursive))
	}
}

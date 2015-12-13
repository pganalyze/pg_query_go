// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeSubselect) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RangeSubselect")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx, "Alias")
	}

	ctx.WriteString(strconv.FormatBool(node.Lateral))

	if node.Subquery != nil {
		node.Subquery.Fingerprint(ctx, "Subquery")
	}
}

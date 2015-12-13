// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeSubselect) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RangeSubselect")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Lateral))

	if node.Subquery != nil {
		node.Subquery.Fingerprint(ctx)
	}
}

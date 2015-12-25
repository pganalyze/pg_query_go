// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeFunction) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RangeFunction")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx, "Alias")
	}

	node.Coldeflist.Fingerprint(ctx, "Coldeflist")
	node.Functions.Fingerprint(ctx, "Functions")
	ctx.WriteString(strconv.FormatBool(node.IsRowsfrom))
	ctx.WriteString(strconv.FormatBool(node.Lateral))
	ctx.WriteString(strconv.FormatBool(node.Ordinality))
}

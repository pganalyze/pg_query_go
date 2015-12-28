// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeFunction) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RangeFunction")

	if node.Alias != nil {
		ctx.WriteString("alias")
		node.Alias.Fingerprint(ctx, "Alias")
	}

	if len(node.Coldeflist.Items) > 0 {
		ctx.WriteString("coldeflist")
		node.Coldeflist.Fingerprint(ctx, "Coldeflist")
	}

	if len(node.Functions.Items) > 0 {
		ctx.WriteString("functions")
		node.Functions.Fingerprint(ctx, "Functions")
	}

	if node.IsRowsfrom {
		ctx.WriteString("is_rowsfrom")
		ctx.WriteString(strconv.FormatBool(node.IsRowsfrom))
	}

	if node.Lateral {
		ctx.WriteString("lateral")
		ctx.WriteString(strconv.FormatBool(node.Lateral))
	}

	if node.Ordinality {
		ctx.WriteString("ordinality")
		ctx.WriteString(strconv.FormatBool(node.Ordinality))
	}
}

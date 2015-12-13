// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeFunction) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RangeFunction")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx, "Alias")
	}

	for _, subNode := range node.Coldeflist {
		subNode.Fingerprint(ctx, "Coldeflist")
	}

	for _, subNode := range node.Functions {
		subNode.Fingerprint(ctx, "Functions")
	}

	ctx.WriteString(strconv.FormatBool(node.IsRowsfrom))
	ctx.WriteString(strconv.FormatBool(node.Lateral))
	ctx.WriteString(strconv.FormatBool(node.Ordinality))
}

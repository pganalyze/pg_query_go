// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeFunction) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RANGEFUNCTION")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx)
	}

	for _, subNode := range node.Coldeflist {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Functions {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.IsRowsfrom))
	ctx.WriteString(strconv.FormatBool(node.Lateral))
	ctx.WriteString(strconv.FormatBool(node.Ordinality))
}

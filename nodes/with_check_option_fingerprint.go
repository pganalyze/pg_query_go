// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node WithCheckOption) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("WithCheckOption")

	if node.Cascaded {
		ctx.WriteString("cascaded")
		ctx.WriteString(strconv.FormatBool(node.Cascaded))
	}

	if node.Qual != nil {
		ctx.WriteString("qual")
		node.Qual.Fingerprint(ctx, "Qual")
	}

	if node.Viewname != nil {
		ctx.WriteString("viewname")
		ctx.WriteString(*node.Viewname)
	}
}

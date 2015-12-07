// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node WithCheckOption) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("WITHCHECKOPTION")
	ctx.WriteString(strconv.FormatBool(node.Cascaded))

	if node.Qual != nil {
		node.Qual.Fingerprint(ctx)
	}

	if node.Viewname != nil {
		ctx.WriteString(*node.Viewname)
	}
}

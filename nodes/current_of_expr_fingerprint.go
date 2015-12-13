// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CurrentOfExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CurrentOfExpr")

	if node.CursorName != nil {
		ctx.WriteString(*node.CursorName)
	}

	ctx.WriteString(strconv.Itoa(int(node.CursorParam)))
	ctx.WriteString(strconv.Itoa(int(node.Cvarno)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

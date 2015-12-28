// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CurrentOfExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CurrentOfExpr")

	if node.CursorName != nil {
		ctx.WriteString("cursor_name")
		ctx.WriteString(*node.CursorName)
	}

	if node.CursorParam != 0 {
		ctx.WriteString("cursor_param")
		ctx.WriteString(strconv.Itoa(int(node.CursorParam)))
	}

	if node.Cvarno != 0 {
		ctx.WriteString("cvarno")
		ctx.WriteString(strconv.Itoa(int(node.Cvarno)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

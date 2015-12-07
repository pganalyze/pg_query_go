// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TargetEntry) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("TARGETENTRY")

	if node.Expr != nil {
		node.Expr.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Resjunk))

	if node.Resname != nil {
		ctx.WriteString(*node.Resname)
	}
}

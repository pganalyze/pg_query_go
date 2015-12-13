// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CaseTestExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CaseTestExpr")
	ctx.WriteString(strconv.Itoa(int(node.Collation)))
	ctx.WriteString(strconv.Itoa(int(node.TypeId)))
	ctx.WriteString(strconv.Itoa(int(node.TypeMod)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}

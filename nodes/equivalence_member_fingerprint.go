// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node EquivalenceMember) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("EQUIVALENCEMEMBER")

	if node.EmExpr != nil {
		node.EmExpr.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.EmIsChild))
	ctx.WriteString(strconv.FormatBool(node.EmIsConst))
}

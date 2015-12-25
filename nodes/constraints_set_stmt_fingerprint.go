// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ConstraintsSetStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ConstraintsSetStmt")
	node.Constraints.Fingerprint(ctx, "Constraints")
	ctx.WriteString(strconv.FormatBool(node.Deferred))
}

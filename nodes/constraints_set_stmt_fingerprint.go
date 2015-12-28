// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ConstraintsSetStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ConstraintsSetStmt")
	if len(node.Constraints.Items) > 0 {
		ctx.WriteString("constraints")
		node.Constraints.Fingerprint(ctx, "Constraints")
	}

	if node.Deferred {
		ctx.WriteString("deferred")
		ctx.WriteString(strconv.FormatBool(node.Deferred))
	}
}

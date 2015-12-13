// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ConstraintsSetStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ConstraintsSetStmt")

	for _, subNode := range node.Constraints {
		subNode.Fingerprint(ctx, "Constraints")
	}

	ctx.WriteString(strconv.FormatBool(node.Deferred))
}

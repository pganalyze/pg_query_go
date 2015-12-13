// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ConstraintsSetStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ConstraintsSetStmt")

	for _, subNode := range node.Constraints {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Deferred))
}

// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node LockStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("LOCK")
	ctx.WriteString(strconv.FormatBool(node.Nowait))

	for _, subNode := range node.Relations {
		subNode.Fingerprint(ctx)
	}
}

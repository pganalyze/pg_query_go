// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node LockStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("LockStmt")
	ctx.WriteString(strconv.Itoa(int(node.Mode)))
	ctx.WriteString(strconv.FormatBool(node.Nowait))

	for _, subNode := range node.Relations {
		subNode.Fingerprint(ctx)
	}
}

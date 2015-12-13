// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DiscardStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("DiscardStmt")
	ctx.WriteString(strconv.Itoa(int(node.Target)))
}

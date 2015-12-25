// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node LockStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("LockStmt")
	ctx.WriteString(strconv.Itoa(int(node.Mode)))
	ctx.WriteString(strconv.FormatBool(node.Nowait))
	node.Relations.Fingerprint(ctx, "Relations")
}

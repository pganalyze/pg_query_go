// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TruncateStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("TruncateStmt")
	ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	node.Relations.Fingerprint(ctx, "Relations")
	ctx.WriteString(strconv.FormatBool(node.RestartSeqs))
}

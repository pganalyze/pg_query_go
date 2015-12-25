// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DropStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DropStmt")
	node.Arguments.Fingerprint(ctx, "Arguments")
	ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	ctx.WriteString(strconv.FormatBool(node.Concurrent))
	ctx.WriteString(strconv.FormatBool(node.MissingOk))
	node.Objects.Fingerprint(ctx, "Objects")
	ctx.WriteString(strconv.Itoa(int(node.RemoveType)))
}

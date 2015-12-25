// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TransactionStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("TransactionStmt")

	if node.Gid != nil {
		ctx.WriteString(*node.Gid)
	}

	ctx.WriteString(strconv.Itoa(int(node.Kind)))
	node.Options.Fingerprint(ctx, "Options")
}

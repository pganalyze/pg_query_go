// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TransactionStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("TransactionStmt")

	if node.Gid != nil {
		ctx.WriteString("gid")
		ctx.WriteString(*node.Gid)
	}

	if int(node.Kind) != 0 {
		ctx.WriteString("kind")
		ctx.WriteString(strconv.Itoa(int(node.Kind)))
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}
}

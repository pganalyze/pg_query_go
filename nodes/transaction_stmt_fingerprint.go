// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TransactionStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("TRANSACTION")

	if node.Gid != nil {
		ctx.WriteString(*node.Gid)
	}

	ctx.WriteString(strconv.Itoa(int(node.Kind)))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}

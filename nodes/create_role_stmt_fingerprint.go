// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateRoleStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CREATEROLESTMT")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Role != nil {
		ctx.WriteString(*node.Role)
	}

	ctx.WriteString(strconv.Itoa(int(node.StmtType)))
}

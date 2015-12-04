// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node GrantRoleStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "GRANTROLESTMT")
	io.WriteString(ctx.hash, strconv.FormatBool(node.AdminOpt))
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Behavior)))

	for _, subNode := range node.GrantedRoles {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.GranteeRoles {
		subNode.Fingerprint(ctx)
	}

	if node.Grantor != nil {
		io.WriteString(ctx.hash, *node.Grantor)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.IsGrant))
}

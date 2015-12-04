// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node DropRoleStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DROPROLESTMT")
	io.WriteString(ctx.hash, strconv.FormatBool(node.MissingOk))

	for _, subNode := range node.Roles {
		subNode.Fingerprint(ctx)
	}
}

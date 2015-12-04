// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node DropOwnedStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DROPOWNEDSTMT")
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Behavior)))

	for _, subNode := range node.Roles {
		subNode.Fingerprint(ctx)
	}
}

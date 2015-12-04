// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node AlterTableMoveAllStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTERTABLEMOVEALLSTMT")

	if node.NewTablespacename != nil {
		io.WriteString(ctx.hash, *node.NewTablespacename)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Nowait))
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Objtype)))

	if node.OrigTablespacename != nil {
		io.WriteString(ctx.hash, *node.OrigTablespacename)
	}

	for _, subNode := range node.Roles {
		subNode.Fingerprint(ctx)
	}
}

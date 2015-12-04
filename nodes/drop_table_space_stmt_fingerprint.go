// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node DropTableSpaceStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DROPTABLESPACESTMT")
	io.WriteString(ctx.hash, strconv.FormatBool(node.MissingOk))

	if node.Tablespacename != nil {
		io.WriteString(ctx.hash, *node.Tablespacename)
	}
}

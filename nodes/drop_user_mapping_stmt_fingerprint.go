// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node DropUserMappingStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DROPUSERMAPPINGSTMT")
	io.WriteString(ctx.hash, strconv.FormatBool(node.MissingOk))

	if node.Servername != nil {
		io.WriteString(ctx.hash, *node.Servername)
	}

	if node.Username != nil {
		io.WriteString(ctx.hash, *node.Username)
	}
}

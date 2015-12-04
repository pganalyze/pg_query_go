// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node AlterTableSpaceOptionsStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTERTABLESPACEOPTIONSSTMT")
	io.WriteString(ctx.hash, strconv.FormatBool(node.IsReset))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Tablespacename != nil {
		io.WriteString(ctx.hash, *node.Tablespacename)
	}
}

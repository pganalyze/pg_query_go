// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node AlterOpFamilyStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTEROPFAMILYSTMT")

	if node.Amname != nil {
		io.WriteString(ctx.hash, *node.Amname)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.IsDrop))

	for _, subNode := range node.Items {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Opfamilyname {
		subNode.Fingerprint(ctx)
	}
}

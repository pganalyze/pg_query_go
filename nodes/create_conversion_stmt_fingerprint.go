// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node CreateConversionStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATECONVERSIONSTMT")

	for _, subNode := range node.ConversionName {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Def))

	if node.ForEncodingName != nil {
		io.WriteString(ctx.hash, *node.ForEncodingName)
	}

	for _, subNode := range node.FuncName {
		subNode.Fingerprint(ctx)
	}

	if node.ToEncodingName != nil {
		io.WriteString(ctx.hash, *node.ToEncodingName)
	}
}

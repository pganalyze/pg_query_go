// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node InlineCodeBlock) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "INLINECODEBLOCK")
	io.WriteString(ctx.hash, strconv.FormatBool(node.LangIsTrusted))

	if node.SourceText != nil {
		io.WriteString(ctx.hash, *node.SourceText)
	}
}

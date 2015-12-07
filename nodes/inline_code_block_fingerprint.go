// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node InlineCodeBlock) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("INLINECODEBLOCK")
	ctx.WriteString(strconv.FormatBool(node.LangIsTrusted))

	if node.SourceText != nil {
		ctx.WriteString(*node.SourceText)
	}
}

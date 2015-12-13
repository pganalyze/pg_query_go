// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node InlineCodeBlock) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("InlineCodeBlock")
	ctx.WriteString(strconv.FormatBool(node.LangIsTrusted))
	ctx.WriteString(strconv.Itoa(int(node.LangOid)))

	if node.SourceText != nil {
		ctx.WriteString(*node.SourceText)
	}
}

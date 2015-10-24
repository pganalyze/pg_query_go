// Auto-generated - DO NOT EDIT

package pg_query

type InlineCodeBlock struct {
	SourceText    *string `json:"source_text"`   /* source text of anonymous code block */
	LangOid       Oid     `json:"langOid"`       /* OID of selected language */
	LangIsTrusted bool    `json:"langIsTrusted"` /* trusted property of the language */
}

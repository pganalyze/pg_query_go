// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type InlineCodeBlock struct {
	SourceText    *string `json:"source_text"`   /* source text of anonymous code block */
	LangOid       Oid     `json:"langOid"`       /* OID of selected language */
	LangIsTrusted bool    `json:"langIsTrusted"` /* trusted property of the language */
}

func (node InlineCodeBlock) MarshalJSON() ([]byte, error) {
	type InlineCodeBlockMarshalAlias InlineCodeBlock
	return json.Marshal(map[string]interface{}{
		"INLINECODEBLOCK": (*InlineCodeBlockMarshalAlias)(&node),
	})
}

func (node *InlineCodeBlock) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

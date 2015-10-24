// Auto-generated - DO NOT EDIT

package pg_query

type XmlSerialize struct {
	Xmloption XmlOptionType `json:"xmloption"` /* DOCUMENT or CONTENT */
	Expr      Node          `json:"expr"`
	TypeName  *TypeName     `json:"typeName"`
	Location  int           `json:"location"` /* token location, or -1 if unknown */
}

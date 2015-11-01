// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type TypeName struct {
	Names       []Node `json:"names"`       /* qualified name (list of Value strings) */
	TypeOid     Oid    `json:"typeOid"`     /* type identified by OID */
	Setof       bool   `json:"setof"`       /* is a set? */
	PctType     bool   `json:"pct_type"`    /* %TYPE specified? */
	Typmods     []Node `json:"typmods"`     /* type modifier expression(s) */
	Typemod     int32  `json:"typemod"`     /* prespecified type modifier */
	ArrayBounds []Node `json:"arrayBounds"` /* array bounds */
	Location    int    `json:"location"`    /* token location, or -1 if unknown */
}

func (node TypeName) MarshalJSON() ([]byte, error) {
	type TypeNameMarshalAlias TypeName
	return json.Marshal(map[string]interface{}{
		"TYPENAME": (*TypeNameMarshalAlias)(&node),
	})
}

func (node *TypeName) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

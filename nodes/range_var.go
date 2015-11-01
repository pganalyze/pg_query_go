// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type RangeVar struct {
	Catalogname *string   `json:"catalogname"` /* the catalog (database) name, or NULL */
	Schemaname  *string   `json:"schemaname"`  /* the schema name, or NULL */
	Relname     *string   `json:"relname"`     /* the relation/sequence name */
	InhOpt      InhOption `json:"inhOpt"`      /* expand rel by inheritance? recursively act
	 * on children? */
	Relpersistence byte   `json:"relpersistence"` /* see RELPERSISTENCE_* in pg_class.h */
	Alias          *Alias `json:"alias"`          /* table alias & optional column aliases */
	Location       int    `json:"location"`       /* token location, or -1 if unknown */
}

func (node RangeVar) MarshalJSON() ([]byte, error) {
	type RangeVarMarshalAlias RangeVar
	return json.Marshal(map[string]interface{}{
		"RANGEVAR": (*RangeVarMarshalAlias)(&node),
	})
}

func (node *RangeVar) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

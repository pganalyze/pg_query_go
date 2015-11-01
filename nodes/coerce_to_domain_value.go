// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CoerceToDomainValue struct {
	Xpr       Expr  `json:"xpr"`
	TypeId    Oid   `json:"typeId"`    /* type for substituted value */
	TypeMod   int32 `json:"typeMod"`   /* typemod for substituted value */
	Collation Oid   `json:"collation"` /* collation for the substituted value */
	Location  int   `json:"location"`  /* token location, or -1 if unknown */
}

func (node CoerceToDomainValue) MarshalJSON() ([]byte, error) {
	type CoerceToDomainValueMarshalAlias CoerceToDomainValue
	return json.Marshal(map[string]interface{}{
		"COERCETODOMAINVALUE": (*CoerceToDomainValueMarshalAlias)(&node),
	})
}

func (node *CoerceToDomainValue) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

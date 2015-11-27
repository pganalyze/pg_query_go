// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["xpr"] != nil {
		err = json.Unmarshal(fields["xpr"], &node.Xpr)
		if err != nil {
			return
		}
	}

	if fields["typeId"] != nil {
		err = json.Unmarshal(fields["typeId"], &node.TypeId)
		if err != nil {
			return
		}
	}

	if fields["typeMod"] != nil {
		err = json.Unmarshal(fields["typeMod"], &node.TypeMod)
		if err != nil {
			return
		}
	}

	if fields["collation"] != nil {
		err = json.Unmarshal(fields["collation"], &node.Collation)
		if err != nil {
			return
		}
	}

	if fields["location"] != nil {
		err = json.Unmarshal(fields["location"], &node.Location)
		if err != nil {
			return
		}
	}

	return
}

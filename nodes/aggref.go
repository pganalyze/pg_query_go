// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Aggref struct {
	Xpr           Expr   `json:"xpr"`
	Aggfnoid      Oid    `json:"aggfnoid"`      /* pg_proc Oid of the aggregate */
	Aggtype       Oid    `json:"aggtype"`       /* type Oid of result of the aggregate */
	Aggcollid     Oid    `json:"aggcollid"`     /* OID of collation of result */
	Inputcollid   Oid    `json:"inputcollid"`   /* OID of collation that function should use */
	Aggdirectargs []Node `json:"aggdirectargs"` /* direct arguments, if an ordered-set agg */
	Args          []Node `json:"args"`          /* aggregated arguments and sort expressions */
	Aggorder      []Node `json:"aggorder"`      /* ORDER BY (list of SortGroupClause) */
	Aggdistinct   []Node `json:"aggdistinct"`   /* DISTINCT (list of SortGroupClause) */
	Aggfilter     *Expr  `json:"aggfilter"`     /* FILTER expression, if any */
	Aggstar       bool   `json:"aggstar"`       /* TRUE if argument list was really '*' */
	Aggvariadic   bool   `json:"aggvariadic"`   /* true if variadic arguments have been
	 * combined into an array last argument */

	Aggkind     byte  `json:"aggkind"`     /* aggregate kind (see pg_aggregate.h) */
	Agglevelsup Index `json:"agglevelsup"` /* > 0 if agg belongs to outer query */
	Location    int   `json:"location"`    /* token location, or -1 if unknown */
}

func (node Aggref) MarshalJSON() ([]byte, error) {
	type AggrefMarshalAlias Aggref
	return json.Marshal(map[string]interface{}{
		"AGGREF": (*AggrefMarshalAlias)(&node),
	})
}

func (node *Aggref) UnmarshalJSON(input []byte) (err error) {
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

	if fields["aggfnoid"] != nil {
		err = json.Unmarshal(fields["aggfnoid"], &node.Aggfnoid)
		if err != nil {
			return
		}
	}

	if fields["aggtype"] != nil {
		err = json.Unmarshal(fields["aggtype"], &node.Aggtype)
		if err != nil {
			return
		}
	}

	if fields["aggcollid"] != nil {
		err = json.Unmarshal(fields["aggcollid"], &node.Aggcollid)
		if err != nil {
			return
		}
	}

	if fields["inputcollid"] != nil {
		err = json.Unmarshal(fields["inputcollid"], &node.Inputcollid)
		if err != nil {
			return
		}
	}

	if fields["aggdirectargs"] != nil {
		node.Aggdirectargs, err = UnmarshalNodeArrayJSON(fields["aggdirectargs"])
		if err != nil {
			return
		}
	}

	if fields["args"] != nil {
		node.Args, err = UnmarshalNodeArrayJSON(fields["args"])
		if err != nil {
			return
		}
	}

	if fields["aggorder"] != nil {
		node.Aggorder, err = UnmarshalNodeArrayJSON(fields["aggorder"])
		if err != nil {
			return
		}
	}

	if fields["aggdistinct"] != nil {
		node.Aggdistinct, err = UnmarshalNodeArrayJSON(fields["aggdistinct"])
		if err != nil {
			return
		}
	}

	if fields["aggfilter"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["aggfilter"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Aggfilter = &val
		}
	}

	if fields["aggstar"] != nil {
		err = json.Unmarshal(fields["aggstar"], &node.Aggstar)
		if err != nil {
			return
		}
	}

	if fields["aggvariadic"] != nil {
		err = json.Unmarshal(fields["aggvariadic"], &node.Aggvariadic)
		if err != nil {
			return
		}
	}

	if fields["aggkind"] != nil {
		var strVal string
		err = json.Unmarshal(fields["aggkind"], &strVal)
		node.Aggkind = strVal[0]
		if err != nil {
			return
		}
	}

	if fields["agglevelsup"] != nil {
		err = json.Unmarshal(fields["agglevelsup"], &node.Agglevelsup)
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

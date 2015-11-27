// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type RangeTblFunction struct {
	Funcexpr     Node `json:"funcexpr"`     /* expression tree for func call */
	Funccolcount int  `json:"funccolcount"` /* number of columns it contributes to RTE */

	/* These fields record the contents of a column definition list, if any: */
	Funccolnames      []Node `json:"funccolnames"`      /* column names (list of String) */
	Funccoltypes      []Node `json:"funccoltypes"`      /* OID list of column type OIDs */
	Funccoltypmods    []Node `json:"funccoltypmods"`    /* integer list of column typmods */
	Funccolcollations []Node `json:"funccolcollations"` /* OID list of column collation OIDs */

	/* This is set during planning for use by the executor: */
	Funcparams []uint32 `json:"funcparams"` /* PARAM_EXEC Param IDs affecting this func */
}

func (node RangeTblFunction) MarshalJSON() ([]byte, error) {
	type RangeTblFunctionMarshalAlias RangeTblFunction
	return json.Marshal(map[string]interface{}{
		"RANGETBLFUNCTION": (*RangeTblFunctionMarshalAlias)(&node),
	})
}

func (node *RangeTblFunction) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["funcexpr"] != nil {
		node.Funcexpr, err = UnmarshalNodeJSON(fields["funcexpr"])
		if err != nil {
			return
		}
	}

	if fields["funccolcount"] != nil {
		err = json.Unmarshal(fields["funccolcount"], &node.Funccolcount)
		if err != nil {
			return
		}
	}

	if fields["funccolnames"] != nil {
		node.Funccolnames, err = UnmarshalNodeArrayJSON(fields["funccolnames"])
		if err != nil {
			return
		}
	}

	if fields["funccoltypes"] != nil {
		node.Funccoltypes, err = UnmarshalNodeArrayJSON(fields["funccoltypes"])
		if err != nil {
			return
		}
	}

	if fields["funccoltypmods"] != nil {
		node.Funccoltypmods, err = UnmarshalNodeArrayJSON(fields["funccoltypmods"])
		if err != nil {
			return
		}
	}

	if fields["funccolcollations"] != nil {
		node.Funccolcollations, err = UnmarshalNodeArrayJSON(fields["funccolcollations"])
		if err != nil {
			return
		}
	}

	if fields["funcparams"] != nil {
		err = json.Unmarshal(fields["funcparams"], &node.Funcparams)
		if err != nil {
			return
		}
	}

	return
}

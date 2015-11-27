// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type WindowAgg struct {
	Plan          Plan        `json:"plan"`
	Winref        Index       `json:"winref"`        /* ID referenced by window functions */
	PartNumCols   int         `json:"partNumCols"`   /* number of columns in partition clause */
	PartColIdx    *AttrNumber `json:"partColIdx"`    /* their indexes in the target list */
	PartOperators *Oid        `json:"partOperators"` /* equality operators for partition columns */
	OrdNumCols    int         `json:"ordNumCols"`    /* number of columns in ordering clause */
	OrdColIdx     *AttrNumber `json:"ordColIdx"`     /* their indexes in the target list */
	OrdOperators  *Oid        `json:"ordOperators"`  /* equality operators for ordering columns */
	FrameOptions  int         `json:"frameOptions"`  /* frame_clause options, see WindowDef */
	StartOffset   Node        `json:"startOffset"`   /* expression for starting bound, if any */
	EndOffset     Node        `json:"endOffset"`     /* expression for ending bound, if any */
}

func (node WindowAgg) MarshalJSON() ([]byte, error) {
	type WindowAggMarshalAlias WindowAgg
	return json.Marshal(map[string]interface{}{
		"WINDOWAGG": (*WindowAggMarshalAlias)(&node),
	})
}

func (node *WindowAgg) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["plan"] != nil {
		err = json.Unmarshal(fields["plan"], &node.Plan)
		if err != nil {
			return
		}
	}

	if fields["winref"] != nil {
		err = json.Unmarshal(fields["winref"], &node.Winref)
		if err != nil {
			return
		}
	}

	if fields["partNumCols"] != nil {
		err = json.Unmarshal(fields["partNumCols"], &node.PartNumCols)
		if err != nil {
			return
		}
	}

	if fields["partColIdx"] != nil {
		err = json.Unmarshal(fields["partColIdx"], &node.PartColIdx)
		if err != nil {
			return
		}
	}

	if fields["partOperators"] != nil {
		err = json.Unmarshal(fields["partOperators"], &node.PartOperators)
		if err != nil {
			return
		}
	}

	if fields["ordNumCols"] != nil {
		err = json.Unmarshal(fields["ordNumCols"], &node.OrdNumCols)
		if err != nil {
			return
		}
	}

	if fields["ordColIdx"] != nil {
		err = json.Unmarshal(fields["ordColIdx"], &node.OrdColIdx)
		if err != nil {
			return
		}
	}

	if fields["ordOperators"] != nil {
		err = json.Unmarshal(fields["ordOperators"], &node.OrdOperators)
		if err != nil {
			return
		}
	}

	if fields["frameOptions"] != nil {
		err = json.Unmarshal(fields["frameOptions"], &node.FrameOptions)
		if err != nil {
			return
		}
	}

	if fields["startOffset"] != nil {
		node.StartOffset, err = UnmarshalNodeJSON(fields["startOffset"])
		if err != nil {
			return
		}
	}

	if fields["endOffset"] != nil {
		node.EndOffset, err = UnmarshalNodeJSON(fields["endOffset"])
		if err != nil {
			return
		}
	}

	return
}

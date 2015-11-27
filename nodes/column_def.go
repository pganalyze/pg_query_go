// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ColumnDef struct {
	Colname       *string        `json:"colname"`        /* name of column */
	TypeName      *TypeName      `json:"typeName"`       /* type of column */
	Inhcount      int            `json:"inhcount"`       /* number of times column is inherited */
	IsLocal       bool           `json:"is_local"`       /* column has local (non-inherited) def'n */
	IsNotNull     bool           `json:"is_not_null"`    /* NOT NULL constraint specified? */
	IsFromType    bool           `json:"is_from_type"`   /* column definition came from table type */
	Storage       byte           `json:"storage"`        /* attstorage setting, or 0 for default */
	RawDefault    Node           `json:"raw_default"`    /* default value (untransformed parse tree) */
	CookedDefault Node           `json:"cooked_default"` /* default value (transformed expr tree) */
	CollClause    *CollateClause `json:"collClause"`     /* untransformed COLLATE spec, if any */
	CollOid       Oid            `json:"collOid"`        /* collation OID (InvalidOid if not set) */
	Constraints   []Node         `json:"constraints"`    /* other constraints on column */
	Fdwoptions    []Node         `json:"fdwoptions"`     /* per-column FDW options */
	Location      int            `json:"location"`       /* parse location, or -1 if none/unknown */
}

func (node ColumnDef) MarshalJSON() ([]byte, error) {
	type ColumnDefMarshalAlias ColumnDef
	return json.Marshal(map[string]interface{}{
		"COLUMNDEF": (*ColumnDefMarshalAlias)(&node),
	})
}

func (node *ColumnDef) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["colname"] != nil {
		err = json.Unmarshal(fields["colname"], &node.Colname)
		if err != nil {
			return
		}
	}

	if fields["typeName"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["typeName"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(TypeName)
			node.TypeName = &val
		}
	}

	if fields["inhcount"] != nil {
		err = json.Unmarshal(fields["inhcount"], &node.Inhcount)
		if err != nil {
			return
		}
	}

	if fields["is_local"] != nil {
		err = json.Unmarshal(fields["is_local"], &node.IsLocal)
		if err != nil {
			return
		}
	}

	if fields["is_not_null"] != nil {
		err = json.Unmarshal(fields["is_not_null"], &node.IsNotNull)
		if err != nil {
			return
		}
	}

	if fields["is_from_type"] != nil {
		err = json.Unmarshal(fields["is_from_type"], &node.IsFromType)
		if err != nil {
			return
		}
	}

	if fields["storage"] != nil {
		var strVal string
		err = json.Unmarshal(fields["storage"], &strVal)
		node.Storage = strVal[0]
		if err != nil {
			return
		}
	}

	if fields["raw_default"] != nil {
		node.RawDefault, err = UnmarshalNodeJSON(fields["raw_default"])
		if err != nil {
			return
		}
	}

	if fields["cooked_default"] != nil {
		node.CookedDefault, err = UnmarshalNodeJSON(fields["cooked_default"])
		if err != nil {
			return
		}
	}

	if fields["collClause"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["collClause"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(CollateClause)
			node.CollClause = &val
		}
	}

	if fields["collOid"] != nil {
		err = json.Unmarshal(fields["collOid"], &node.CollOid)
		if err != nil {
			return
		}
	}

	if fields["constraints"] != nil {
		node.Constraints, err = UnmarshalNodeArrayJSON(fields["constraints"])
		if err != nil {
			return
		}
	}

	if fields["fdwoptions"] != nil {
		node.Fdwoptions, err = UnmarshalNodeArrayJSON(fields["fdwoptions"])
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

// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type TargetEntry struct {
	Xpr             Expr       `json:"xpr"`
	Expr            *Expr      `json:"expr"`            /* expression to evaluate */
	Resno           AttrNumber `json:"resno"`           /* attribute number (see notes above) */
	Resname         *string    `json:"resname"`         /* name of the column (could be NULL) */
	Ressortgroupref Index      `json:"ressortgroupref"` /* nonzero if referenced by a sort/group
	 * clause */

	Resorigtbl Oid        `json:"resorigtbl"` /* OID of column's source table */
	Resorigcol AttrNumber `json:"resorigcol"` /* column's number in source table */
	Resjunk    bool       `json:"resjunk"`    /* set to true to eliminate the attribute from
	 * final target list */

}

func (node TargetEntry) MarshalJSON() ([]byte, error) {
	type TargetEntryMarshalAlias TargetEntry
	return json.Marshal(map[string]interface{}{
		"TARGETENTRY": (*TargetEntryMarshalAlias)(&node),
	})
}

func (node *TargetEntry) UnmarshalJSON(input []byte) (err error) {
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

	if fields["expr"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["expr"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Expr = &val
		}
	}

	if fields["resno"] != nil {
		err = json.Unmarshal(fields["resno"], &node.Resno)
		if err != nil {
			return
		}
	}

	if fields["resname"] != nil {
		err = json.Unmarshal(fields["resname"], &node.Resname)
		if err != nil {
			return
		}
	}

	if fields["ressortgroupref"] != nil {
		err = json.Unmarshal(fields["ressortgroupref"], &node.Ressortgroupref)
		if err != nil {
			return
		}
	}

	if fields["resorigtbl"] != nil {
		err = json.Unmarshal(fields["resorigtbl"], &node.Resorigtbl)
		if err != nil {
			return
		}
	}

	if fields["resorigcol"] != nil {
		err = json.Unmarshal(fields["resorigcol"], &node.Resorigcol)
		if err != nil {
			return
		}
	}

	if fields["resjunk"] != nil {
		err = json.Unmarshal(fields["resjunk"], &node.Resjunk)
		if err != nil {
			return
		}
	}

	return
}

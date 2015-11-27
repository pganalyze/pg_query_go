// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type IntoClause struct {
	Rel            *RangeVar      `json:"rel"`            /* target relation name */
	ColNames       []Node         `json:"colNames"`       /* column names to assign, or NIL */
	Options        []Node         `json:"options"`        /* options from WITH clause */
	OnCommit       OnCommitAction `json:"onCommit"`       /* what do we do at COMMIT? */
	TableSpaceName *string        `json:"tableSpaceName"` /* table space to use, or NULL */
	ViewQuery      Node           `json:"viewQuery"`      /* materialized view's SELECT query */
	SkipData       bool           `json:"skipData"`       /* true for WITH NO DATA */
}

func (node IntoClause) MarshalJSON() ([]byte, error) {
	type IntoClauseMarshalAlias IntoClause
	return json.Marshal(map[string]interface{}{
		"INTOCLAUSE": (*IntoClauseMarshalAlias)(&node),
	})
}

func (node *IntoClause) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["rel"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["rel"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(RangeVar)
			node.Rel = &val
		}
	}

	if fields["colNames"] != nil {
		node.ColNames, err = UnmarshalNodeArrayJSON(fields["colNames"])
		if err != nil {
			return
		}
	}

	if fields["options"] != nil {
		node.Options, err = UnmarshalNodeArrayJSON(fields["options"])
		if err != nil {
			return
		}
	}

	if fields["onCommit"] != nil {
		err = json.Unmarshal(fields["onCommit"], &node.OnCommit)
		if err != nil {
			return
		}
	}

	if fields["tableSpaceName"] != nil {
		err = json.Unmarshal(fields["tableSpaceName"], &node.TableSpaceName)
		if err != nil {
			return
		}
	}

	if fields["viewQuery"] != nil {
		node.ViewQuery, err = UnmarshalNodeJSON(fields["viewQuery"])
		if err != nil {
			return
		}
	}

	if fields["skipData"] != nil {
		err = json.Unmarshal(fields["skipData"], &node.SkipData)
		if err != nil {
			return
		}
	}

	return
}

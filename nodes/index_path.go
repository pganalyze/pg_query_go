// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type IndexPath struct {
	Path             Path          `json:"path"`
	Indexinfo        *IndexOptInfo `json:"indexinfo"`
	Indexclauses     []Node        `json:"indexclauses"`
	Indexquals       []Node        `json:"indexquals"`
	Indexqualcols    []Node        `json:"indexqualcols"`
	Indexorderbys    []Node        `json:"indexorderbys"`
	Indexorderbycols []Node        `json:"indexorderbycols"`
	Indexscandir     ScanDirection `json:"indexscandir"`
	Indextotalcost   Cost          `json:"indextotalcost"`
	Indexselectivity Selectivity   `json:"indexselectivity"`
}

func (node IndexPath) MarshalJSON() ([]byte, error) {
	type IndexPathMarshalAlias IndexPath
	return json.Marshal(map[string]interface{}{
		"INDEXPATH": (*IndexPathMarshalAlias)(&node),
	})
}

func (node *IndexPath) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["path"] != nil {
		err = json.Unmarshal(fields["path"], &node.Path)
		if err != nil {
			return
		}
	}

	if fields["indexinfo"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["indexinfo"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(IndexOptInfo)
			node.Indexinfo = &val
		}
	}

	if fields["indexclauses"] != nil {
		node.Indexclauses, err = UnmarshalNodeArrayJSON(fields["indexclauses"])
		if err != nil {
			return
		}
	}

	if fields["indexquals"] != nil {
		node.Indexquals, err = UnmarshalNodeArrayJSON(fields["indexquals"])
		if err != nil {
			return
		}
	}

	if fields["indexqualcols"] != nil {
		node.Indexqualcols, err = UnmarshalNodeArrayJSON(fields["indexqualcols"])
		if err != nil {
			return
		}
	}

	if fields["indexorderbys"] != nil {
		node.Indexorderbys, err = UnmarshalNodeArrayJSON(fields["indexorderbys"])
		if err != nil {
			return
		}
	}

	if fields["indexorderbycols"] != nil {
		node.Indexorderbycols, err = UnmarshalNodeArrayJSON(fields["indexorderbycols"])
		if err != nil {
			return
		}
	}

	if fields["indexscandir"] != nil {
		err = json.Unmarshal(fields["indexscandir"], &node.Indexscandir)
		if err != nil {
			return
		}
	}

	if fields["indextotalcost"] != nil {
		err = json.Unmarshal(fields["indextotalcost"], &node.Indextotalcost)
		if err != nil {
			return
		}
	}

	if fields["indexselectivity"] != nil {
		err = json.Unmarshal(fields["indexselectivity"], &node.Indexselectivity)
		if err != nil {
			return
		}
	}

	return
}

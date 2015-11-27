// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type IndexOnlyScan struct {
	Scan          Scan          `json:"scan"`
	Indexid       Oid           `json:"indexid"`       /* OID of index to scan */
	Indexqual     []Node        `json:"indexqual"`     /* list of index quals (usually OpExprs) */
	Indexorderby  []Node        `json:"indexorderby"`  /* list of index ORDER BY exprs */
	Indextlist    []Node        `json:"indextlist"`    /* TargetEntry list describing index's cols */
	Indexorderdir ScanDirection `json:"indexorderdir"` /* forward or backward or don't care */
}

func (node IndexOnlyScan) MarshalJSON() ([]byte, error) {
	type IndexOnlyScanMarshalAlias IndexOnlyScan
	return json.Marshal(map[string]interface{}{
		"INDEXONLYSCAN": (*IndexOnlyScanMarshalAlias)(&node),
	})
}

func (node *IndexOnlyScan) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["scan"] != nil {
		err = json.Unmarshal(fields["scan"], &node.Scan)
		if err != nil {
			return
		}
	}

	if fields["indexid"] != nil {
		err = json.Unmarshal(fields["indexid"], &node.Indexid)
		if err != nil {
			return
		}
	}

	if fields["indexqual"] != nil {
		node.Indexqual, err = UnmarshalNodeArrayJSON(fields["indexqual"])
		if err != nil {
			return
		}
	}

	if fields["indexorderby"] != nil {
		node.Indexorderby, err = UnmarshalNodeArrayJSON(fields["indexorderby"])
		if err != nil {
			return
		}
	}

	if fields["indextlist"] != nil {
		node.Indextlist, err = UnmarshalNodeArrayJSON(fields["indextlist"])
		if err != nil {
			return
		}
	}

	if fields["indexorderdir"] != nil {
		err = json.Unmarshal(fields["indexorderdir"], &node.Indexorderdir)
		if err != nil {
			return
		}
	}

	return
}

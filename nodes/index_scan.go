// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type IndexScan struct {
	Scan             Scan          `json:"scan"`
	Indexid          Oid           `json:"indexid"`          /* OID of index to scan */
	Indexqual        []Node        `json:"indexqual"`        /* list of index quals (usually OpExprs) */
	Indexqualorig    []Node        `json:"indexqualorig"`    /* the same in original form */
	Indexorderby     []Node        `json:"indexorderby"`     /* list of index ORDER BY exprs */
	Indexorderbyorig []Node        `json:"indexorderbyorig"` /* the same in original form */
	Indexorderdir    ScanDirection `json:"indexorderdir"`    /* forward or backward or don't care */
}

func (node IndexScan) MarshalJSON() ([]byte, error) {
	type IndexScanMarshalAlias IndexScan
	return json.Marshal(map[string]interface{}{
		"INDEXSCAN": (*IndexScanMarshalAlias)(&node),
	})
}

func (node *IndexScan) UnmarshalJSON(input []byte) (err error) {
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

	if fields["indexqualorig"] != nil {
		node.Indexqualorig, err = UnmarshalNodeArrayJSON(fields["indexqualorig"])
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

	if fields["indexorderbyorig"] != nil {
		node.Indexorderbyorig, err = UnmarshalNodeArrayJSON(fields["indexorderbyorig"])
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

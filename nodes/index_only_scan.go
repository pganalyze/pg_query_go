// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 *		index-only scan node
 *
 * IndexOnlyScan is very similar to IndexScan, but it specifies an
 * index-only scan, in which the data comes from the index not the heap.
 * Because of this, *all* Vars in the plan node's targetlist, qual, and
 * index expressions reference index columns and have varno = INDEX_VAR.
 * Hence we do not need separate indexqualorig and indexorderbyorig lists,
 * since their contents would be equivalent to indexqual and indexorderby.
 *
 * To help EXPLAIN interpret the index Vars for display, we provide
 * indextlist, which represents the contents of the index as a targetlist
 * with one TLE per index column.  Vars appearing in this list reference
 * the base table, and this is the only field in the plan node that may
 * contain such Vars.
 * ----------------
 */
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

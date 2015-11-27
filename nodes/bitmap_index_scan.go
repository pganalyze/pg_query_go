// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 *		bitmap index scan node
 *
 * BitmapIndexScan delivers a bitmap of potential tuple locations;
 * it does not access the heap itself.  The bitmap is used by an
 * ancestor BitmapHeapScan node, possibly after passing through
 * intermediate BitmapAnd and/or BitmapOr nodes to combine it with
 * the results of other BitmapIndexScans.
 *
 * The fields have the same meanings as for IndexScan, except we don't
 * store a direction flag because direction is uninteresting.
 *
 * In a BitmapIndexScan plan node, the targetlist and qual fields are
 * not used and are always NIL.  The indexqualorig field is unused at
 * run time too, but is saved for the benefit of EXPLAIN.
 * ----------------
 */
type BitmapIndexScan struct {
	Scan          Scan   `json:"scan"`
	Indexid       Oid    `json:"indexid"`       /* OID of index to scan */
	Indexqual     []Node `json:"indexqual"`     /* list of index quals (OpExprs) */
	Indexqualorig []Node `json:"indexqualorig"` /* the same in original form */
}

func (node BitmapIndexScan) MarshalJSON() ([]byte, error) {
	type BitmapIndexScanMarshalAlias BitmapIndexScan
	return json.Marshal(map[string]interface{}{
		"BITMAPINDEXSCAN": (*BitmapIndexScanMarshalAlias)(&node),
	})
}

func (node *BitmapIndexScan) UnmarshalJSON(input []byte) (err error) {
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

	return
}

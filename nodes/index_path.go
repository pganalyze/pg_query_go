// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*----------
 * IndexPath represents an index scan over a single index.
 *
 * This struct is used for both regular indexscans and index-only scans;
 * path.pathtype is T_IndexScan or T_IndexOnlyScan to show which is meant.
 *
 * 'indexinfo' is the index to be scanned.
 *
 * 'indexclauses' is a list of index qualification clauses, with implicit
 * AND semantics across the list.  Each clause is a RestrictInfo node from
 * the query's WHERE or JOIN conditions.  An empty list implies a full
 * index scan.
 *
 * 'indexquals' has the same structure as 'indexclauses', but it contains
 * the actual index qual conditions that can be used with the index.
 * In simple cases this is identical to 'indexclauses', but when special
 * indexable operators appear in 'indexclauses', they are replaced by the
 * derived indexscannable conditions in 'indexquals'.
 *
 * 'indexqualcols' is an integer list of index column numbers (zero-based)
 * of the same length as 'indexquals', showing which index column each qual
 * is meant to be used with.  'indexquals' is required to be ordered by
 * index column, so 'indexqualcols' must form a nondecreasing sequence.
 * (The order of multiple quals for the same index column is unspecified.)
 *
 * 'indexorderbys', if not NIL, is a list of ORDER BY expressions that have
 * been found to be usable as ordering operators for an amcanorderbyop index.
 * The list must match the path's pathkeys, ie, one expression per pathkey
 * in the same order.  These are not RestrictInfos, just bare expressions,
 * since they generally won't yield booleans.  Also, unlike the case for
 * quals, it's guaranteed that each expression has the index key on the left
 * side of the operator.
 *
 * 'indexorderbycols' is an integer list of index column numbers (zero-based)
 * of the same length as 'indexorderbys', showing which index column each
 * ORDER BY expression is meant to be used with.  (There is no restriction
 * on which index column each ORDER BY can be used with.)
 *
 * 'indexscandir' is one of:
 *		ForwardScanDirection: forward scan of an ordered index
 *		BackwardScanDirection: backward scan of an ordered index
 *		NoMovementScanDirection: scan of an unordered index, or don't care
 * (The executor doesn't care whether it gets ForwardScanDirection or
 * NoMovementScanDirection for an indexscan, but the planner wants to
 * distinguish ordered from unordered indexes for building pathkeys.)
 *
 * 'indextotalcost' and 'indexselectivity' are saved in the IndexPath so that
 * we need not recompute them when considering using the same index in a
 * bitmap index/heap scan (see BitmapHeapPath).  The costs of the IndexPath
 * itself represent the costs of an IndexScan or IndexOnlyScan plan type.
 *----------
 */
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

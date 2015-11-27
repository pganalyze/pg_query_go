// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 *		index scan node
 *
 * indexqualorig is an implicitly-ANDed list of index qual expressions, each
 * in the same form it appeared in the query WHERE condition.  Each should
 * be of the form (indexkey OP comparisonval) or (comparisonval OP indexkey).
 * The indexkey is a Var or expression referencing column(s) of the index's
 * base table.  The comparisonval might be any expression, but it won't use
 * any columns of the base table.  The expressions are ordered by index
 * column position (but items referencing the same index column can appear
 * in any order).  indexqualorig is used at runtime only if we have to recheck
 * a lossy indexqual.
 *
 * indexqual has the same form, but the expressions have been commuted if
 * necessary to put the indexkeys on the left, and the indexkeys are replaced
 * by Var nodes identifying the index columns (their varno is INDEX_VAR and
 * their varattno is the index column number).
 *
 * indexorderbyorig is similarly the original form of any ORDER BY expressions
 * that are being implemented by the index, while indexorderby is modified to
 * have index column Vars on the left-hand side.  Here, multiple expressions
 * must appear in exactly the ORDER BY order, and this is not necessarily the
 * index column order.  Only the expressions are provided, not the auxiliary
 * sort-order information from the ORDER BY SortGroupClauses; it's assumed
 * that the sort ordering is fully determinable from the top-level operators.
 * indexorderbyorig is unused at run time, but is needed for EXPLAIN.
 * (Note these fields are used for amcanorderbyop cases, not amcanorder cases.)
 *
 * indexorderdir specifies the scan ordering, for indexscans on amcanorder
 * indexes (for other indexes it should be "don't care").
 * ----------------
 */
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

// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 *		ForeignScan node
 *
 * fdw_exprs and fdw_private are both under the control of the foreign-data
 * wrapper, but fdw_exprs is presumed to contain expression trees and will
 * be post-processed accordingly by the planner; fdw_private won't be.
 * Note that everything in both lists must be copiable by copyObject().
 * One way to store an arbitrary blob of bytes is to represent it as a bytea
 * Const.  Usually, though, you'll be better off choosing a representation
 * that can be dumped usefully by nodeToString().
 * ----------------
 */
type ForeignScan struct {
	Scan        Scan   `json:"scan"`
	FdwExprs    []Node `json:"fdw_exprs"`   /* expressions that FDW may evaluate */
	FdwPrivate  []Node `json:"fdw_private"` /* private data for FDW */
	FsSystemCol bool   `json:"fsSystemCol"` /* true if any "system column" is needed */
}

func (node ForeignScan) MarshalJSON() ([]byte, error) {
	type ForeignScanMarshalAlias ForeignScan
	return json.Marshal(map[string]interface{}{
		"FOREIGNSCAN": (*ForeignScanMarshalAlias)(&node),
	})
}

func (node *ForeignScan) UnmarshalJSON(input []byte) (err error) {
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

	if fields["fdw_exprs"] != nil {
		node.FdwExprs, err = UnmarshalNodeArrayJSON(fields["fdw_exprs"])
		if err != nil {
			return
		}
	}

	if fields["fdw_private"] != nil {
		node.FdwPrivate, err = UnmarshalNodeArrayJSON(fields["fdw_private"])
		if err != nil {
			return
		}
	}

	if fields["fsSystemCol"] != nil {
		err = json.Unmarshal(fields["fsSystemCol"], &node.FsSystemCol)
		if err != nil {
			return
		}
	}

	return
}

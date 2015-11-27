// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

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

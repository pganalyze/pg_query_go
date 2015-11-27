// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 *		bitmap sequential scan node
 *
 * This needs a copy of the qual conditions being used by the input index
 * scans because there are various cases where we need to recheck the quals;
 * for example, when the bitmap is lossy about the specific rows on a page
 * that meet the index condition.
 * ----------------
 */
type BitmapHeapScan struct {
	Scan           Scan   `json:"scan"`
	Bitmapqualorig []Node `json:"bitmapqualorig"` /* index quals, in standard expr form */
}

func (node BitmapHeapScan) MarshalJSON() ([]byte, error) {
	type BitmapHeapScanMarshalAlias BitmapHeapScan
	return json.Marshal(map[string]interface{}{
		"BITMAPHEAPSCAN": (*BitmapHeapScanMarshalAlias)(&node),
	})
}

func (node *BitmapHeapScan) UnmarshalJSON(input []byte) (err error) {
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

	if fields["bitmapqualorig"] != nil {
		node.Bitmapqualorig, err = UnmarshalNodeArrayJSON(fields["bitmapqualorig"])
		if err != nil {
			return
		}
	}

	return
}

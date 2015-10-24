// Auto-generated - DO NOT EDIT

package pg_query

type BitmapAndPath struct {
	Path              Path        `json:"path"`
	Bitmapquals       []Node      `json:"bitmapquals"` /* IndexPaths and BitmapOrPaths */
	Bitmapselectivity Selectivity `json:"bitmapselectivity"`
}

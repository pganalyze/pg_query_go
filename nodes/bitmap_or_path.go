// Auto-generated - DO NOT EDIT

package pg_query

type BitmapOrPath struct {
	Path              Path        `json:"path"`
	Bitmapquals       []Node      `json:"bitmapquals"` /* IndexPaths and BitmapAndPaths */
	Bitmapselectivity Selectivity `json:"bitmapselectivity"`
}

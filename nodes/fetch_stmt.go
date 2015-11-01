// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type FetchStmt struct {
	Direction  FetchDirection `json:"direction"`  /* see above */
	HowMany    int64          `json:"howMany"`    /* number of rows, or position argument */
	Portalname *string        `json:"portalname"` /* name of portal (cursor) */
	Ismove     bool           `json:"ismove"`     /* TRUE if MOVE */
}

func (node FetchStmt) MarshalJSON() ([]byte, error) {
	type FetchStmtMarshalAlias FetchStmt
	return json.Marshal(map[string]interface{}{
		"FETCHSTMT": (*FetchStmtMarshalAlias)(&node),
	})
}

func (node *FetchStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

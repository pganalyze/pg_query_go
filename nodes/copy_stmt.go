// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CopyStmt struct {
	Relation *RangeVar `json:"relation"` /* the relation to copy */
	Query    Node      `json:"query"`    /* the SELECT query to copy */
	Attlist  []Node    `json:"attlist"`  /* List of column names (as Strings), or NIL
	 * for all columns */
	IsFrom    bool    `json:"is_from"`    /* TO or FROM */
	IsProgram bool    `json:"is_program"` /* is 'filename' a program to popen? */
	Filename  *string `json:"filename"`   /* filename, or NULL for STDIN/STDOUT */
	Options   []Node  `json:"options"`    /* List of DefElem nodes */
}

func (node CopyStmt) MarshalJSON() ([]byte, error) {
	type CopyStmtMarshalAlias CopyStmt
	return json.Marshal(map[string]interface{}{
		"COPY": (*CopyStmtMarshalAlias)(&node),
	})
}

func (node *CopyStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

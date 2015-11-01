// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterTSConfigurationStmt struct {
	Cfgname []Node `json:"cfgname"` /* qualified name (list of Value strings) */

	/*
	 * dicts will be non-NIL if ADD/ALTER MAPPING was specified. If dicts is
	 * NIL, but tokentype isn't, DROP MAPPING was specified.
	 */
	Tokentype []Node `json:"tokentype"`  /* list of Value strings */
	Dicts     []Node `json:"dicts"`      /* list of list of Value strings */
	Override  bool   `json:"override"`   /* if true - remove old variant */
	Replace   bool   `json:"replace"`    /* if true - replace dictionary by another */
	MissingOk bool   `json:"missing_ok"` /* for DROP - skip error if missing? */
}

func (node AlterTSConfigurationStmt) MarshalJSON() ([]byte, error) {
	type AlterTSConfigurationStmtMarshalAlias AlterTSConfigurationStmt
	return json.Marshal(map[string]interface{}{
		"ALTERTSCONFIGURATIONSTMT": (*AlterTSConfigurationStmtMarshalAlias)(&node),
	})
}

func (node *AlterTSConfigurationStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

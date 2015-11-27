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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["cfgname"] != nil {
		node.Cfgname, err = UnmarshalNodeArrayJSON(fields["cfgname"])
		if err != nil {
			return
		}
	}

	if fields["tokentype"] != nil {
		node.Tokentype, err = UnmarshalNodeArrayJSON(fields["tokentype"])
		if err != nil {
			return
		}
	}

	if fields["dicts"] != nil {
		node.Dicts, err = UnmarshalNodeArrayJSON(fields["dicts"])
		if err != nil {
			return
		}
	}

	if fields["override"] != nil {
		err = json.Unmarshal(fields["override"], &node.Override)
		if err != nil {
			return
		}
	}

	if fields["replace"] != nil {
		err = json.Unmarshal(fields["replace"], &node.Replace)
		if err != nil {
			return
		}
	}

	if fields["missing_ok"] != nil {
		err = json.Unmarshal(fields["missing_ok"], &node.MissingOk)
		if err != nil {
			return
		}
	}

	return
}

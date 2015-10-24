// Auto-generated - DO NOT EDIT

package pg_query

type CreatePLangStmt struct {
	Replace     bool    `json:"replace"`     /* T => replace if already exists */
	Plname      *string `json:"plname"`      /* PL name */
	Plhandler   []Node  `json:"plhandler"`   /* PL call handler function (qual. name) */
	Plinline    []Node  `json:"plinline"`    /* optional inline function (qual. name) */
	Plvalidator []Node  `json:"plvalidator"` /* optional validator function (qual. name) */
	Pltrusted   bool    `json:"pltrusted"`   /* PL is trusted */
}

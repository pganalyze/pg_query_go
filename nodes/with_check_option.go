// Auto-generated - DO NOT EDIT

package pg_query

type WithCheckOption struct {
	Viewname *string `json:"viewname"` /* name of view that specified the WCO */
	Qual     Node    `json:"qual"`     /* constraint qual to check */
	Cascaded bool    `json:"cascaded"` /* true = WITH CASCADED CHECK OPTION */
}

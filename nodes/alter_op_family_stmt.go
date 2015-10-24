// Auto-generated - DO NOT EDIT

package pg_query

type AlterOpFamilyStmt struct {
	Opfamilyname []Node  `json:"opfamilyname"` /* qualified name (list of Value strings) */
	Amname       *string `json:"amname"`       /* name of index AM opfamily is for */
	IsDrop       bool    `json:"isDrop"`       /* ADD or DROP the items? */
	Items        []Node  `json:"items"`        /* List of CreateOpClassItem nodes */
}

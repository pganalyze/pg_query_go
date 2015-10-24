// Auto-generated - DO NOT EDIT

package pg_query

type SubLink struct {
	Xpr         Expr        `json:"xpr"`
	SubLinkType SubLinkType `json:"subLinkType"` /* see above */
	Testexpr    Node        `json:"testexpr"`    /* outer-query test for ALL/ANY/ROWCOMPARE */
	OperName    []Node      `json:"operName"`    /* originally specified operator name */
	Subselect   Node        `json:"subselect"`   /* subselect as Query* or parsetree */
	Location    int         `json:"location"`    /* token location, or -1 if unknown */
}

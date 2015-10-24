// Auto-generated - DO NOT EDIT

package pg_query

type SortBy struct {
	Node        Node        `json:"node"`         /* expression to sort on */
	SortbyDir   SortByDir   `json:"sortby_dir"`   /* ASC/DESC/USING/default */
	SortbyNulls SortByNulls `json:"sortby_nulls"` /* NULLS FIRST/LAST */
	UseOp       []Node      `json:"useOp"`        /* name of op to use, if SORTBY_USING */
	Location    int         `json:"location"`     /* operator location, or -1 if none/unknown */
}

// Auto-generated - DO NOT EDIT

package pg_query

type IndexElem struct {
	Name          *string     `json:"name"`           /* name of attribute to index, or NULL */
	Expr          Node        `json:"expr"`           /* expression to index, or NULL */
	Indexcolname  *string     `json:"indexcolname"`   /* name for index column; NULL = default */
	Collation     []Node      `json:"collation"`      /* name of collation; NIL = default */
	Opclass       []Node      `json:"opclass"`        /* name of desired opclass; NIL = default */
	Ordering      SortByDir   `json:"ordering"`       /* ASC/DESC/default */
	NullsOrdering SortByNulls `json:"nulls_ordering"` /* FIRST/LAST/default */
}

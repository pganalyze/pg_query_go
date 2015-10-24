// Auto-generated - DO NOT EDIT

package pg_query

type CreateOpClassItem struct {
	Itemtype int `json:"itemtype"` /* see codes above */
	/* fields used for an operator or function item: */
	Name        []Node `json:"name"`         /* operator or function name */
	Args        []Node `json:"args"`         /* argument types */
	Number      int    `json:"number"`       /* strategy num or support proc num */
	OrderFamily []Node `json:"order_family"` /* only used for ordering operators */
	ClassArgs   []Node `json:"class_args"`   /* only used for functions */
	/* fields used for a storagetype item: */
	Storedtype *TypeName `json:"storedtype"` /* datatype stored in index */
}

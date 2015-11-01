// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node CreateOpClassItem) MarshalJSON() ([]byte, error) {
	type CreateOpClassItemMarshalAlias CreateOpClassItem
	return json.Marshal(map[string]interface{}{
		"CREATEOPCLASSITEM": (*CreateOpClassItemMarshalAlias)(&node),
	})
}

func (node *CreateOpClassItem) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

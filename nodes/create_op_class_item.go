// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------------
 *		Create Operator Class Statement
 * ----------------------
 */
type CreateOpClassItem struct {
	Itemtype int `json:"itemtype"` /* see codes above */

	/* fields used for an operator or function item: */
	Name        List `json:"name"`         /* operator or function name */
	Args        List `json:"args"`         /* argument types */
	Number      int  `json:"number"`       /* strategy num or support proc num */
	OrderFamily List `json:"order_family"` /* only used for ordering operators */
	ClassArgs   List `json:"class_args"`   /* only used for functions */

	/* fields used for a storagetype item: */
	Storedtype *TypeName `json:"storedtype"` /* datatype stored in index */
}

func (node CreateOpClassItem) MarshalJSON() ([]byte, error) {
	type CreateOpClassItemMarshalAlias CreateOpClassItem
	return json.Marshal(map[string]interface{}{
		"CreateOpClassItem": (*CreateOpClassItemMarshalAlias)(&node),
	})
}

func (node *CreateOpClassItem) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["itemtype"] != nil {
		err = json.Unmarshal(fields["itemtype"], &node.Itemtype)
		if err != nil {
			return
		}
	}

	if fields["name"] != nil {
		node.Name.Items, err = UnmarshalNodeArrayJSON(fields["name"])
		if err != nil {
			return
		}
	}

	if fields["args"] != nil {
		node.Args.Items, err = UnmarshalNodeArrayJSON(fields["args"])
		if err != nil {
			return
		}
	}

	if fields["number"] != nil {
		err = json.Unmarshal(fields["number"], &node.Number)
		if err != nil {
			return
		}
	}

	if fields["order_family"] != nil {
		node.OrderFamily.Items, err = UnmarshalNodeArrayJSON(fields["order_family"])
		if err != nil {
			return
		}
	}

	if fields["class_args"] != nil {
		node.ClassArgs.Items, err = UnmarshalNodeArrayJSON(fields["class_args"])
		if err != nil {
			return
		}
	}

	if fields["storedtype"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["storedtype"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(TypeName)
			node.Storedtype = &val
		}
	}

	return
}

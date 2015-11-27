// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type SubLink struct {
	Xpr         Expr        `json:"xpr"`
	SubLinkType SubLinkType `json:"subLinkType"` /* see above */
	Testexpr    Node        `json:"testexpr"`    /* outer-query test for ALL/ANY/ROWCOMPARE */
	OperName    []Node      `json:"operName"`    /* originally specified operator name */
	Subselect   Node        `json:"subselect"`   /* subselect as Query* or parsetree */
	Location    int         `json:"location"`    /* token location, or -1 if unknown */
}

func (node SubLink) MarshalJSON() ([]byte, error) {
	type SubLinkMarshalAlias SubLink
	return json.Marshal(map[string]interface{}{
		"SUBLINK": (*SubLinkMarshalAlias)(&node),
	})
}

func (node *SubLink) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["xpr"] != nil {
		err = json.Unmarshal(fields["xpr"], &node.Xpr)
		if err != nil {
			return
		}
	}

	if fields["subLinkType"] != nil {
		err = json.Unmarshal(fields["subLinkType"], &node.SubLinkType)
		if err != nil {
			return
		}
	}

	if fields["testexpr"] != nil {
		node.Testexpr, err = UnmarshalNodeJSON(fields["testexpr"])
		if err != nil {
			return
		}
	}

	if fields["operName"] != nil {
		node.OperName, err = UnmarshalNodeArrayJSON(fields["operName"])
		if err != nil {
			return
		}
	}

	if fields["subselect"] != nil {
		node.Subselect, err = UnmarshalNodeJSON(fields["subselect"])
		if err != nil {
			return
		}
	}

	if fields["location"] != nil {
		err = json.Unmarshal(fields["location"], &node.Location)
		if err != nil {
			return
		}
	}

	return
}

// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type JoinExpr struct {
	Jointype    JoinType `json:"jointype"`    /* type of join */
	IsNatural   bool     `json:"isNatural"`   /* Natural join? Will need to shape table */
	Larg        Node     `json:"larg"`        /* left subtree */
	Rarg        Node     `json:"rarg"`        /* right subtree */
	UsingClause []Node   `json:"usingClause"` /* USING clause, if any (list of String) */
	Quals       Node     `json:"quals"`       /* qualifiers on join, if any */
	Alias       *Alias   `json:"alias"`       /* user-written alias clause, if any */
	Rtindex     int      `json:"rtindex"`     /* RT index assigned for join, or 0 */
}

func (node JoinExpr) MarshalJSON() ([]byte, error) {
	type JoinExprMarshalAlias JoinExpr
	return json.Marshal(map[string]interface{}{
		"JOINEXPR": (*JoinExprMarshalAlias)(&node),
	})
}

func (node *JoinExpr) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

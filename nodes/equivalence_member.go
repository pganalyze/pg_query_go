// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type EquivalenceMember struct {
	EmExpr           *Expr    `json:"em_expr"`            /* the expression represented */
	EmRelids         []uint32 `json:"em_relids"`          /* all relids appearing in em_expr */
	EmNullableRelids []uint32 `json:"em_nullable_relids"` /* nullable by lower outer joins */
	EmIsConst        bool     `json:"em_is_const"`        /* expression is pseudoconstant? */
	EmIsChild        bool     `json:"em_is_child"`        /* derived version for a child relation? */
	EmDatatype       Oid      `json:"em_datatype"`        /* the "nominal type" used by the opfamily */
}

func (node EquivalenceMember) MarshalJSON() ([]byte, error) {
	type EquivalenceMemberMarshalAlias EquivalenceMember
	return json.Marshal(map[string]interface{}{
		"EQUIVALENCEMEMBER": (*EquivalenceMemberMarshalAlias)(&node),
	})
}

func (node *EquivalenceMember) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateDomainStmt struct {
	Domainname  []Node         `json:"domainname"`  /* qualified name (list of Value strings) */
	TypeName    *TypeName      `json:"typeName"`    /* the base type */
	CollClause  *CollateClause `json:"collClause"`  /* untransformed COLLATE spec, if any */
	Constraints []Node         `json:"constraints"` /* constraints (list of Constraint nodes) */
}

func (node CreateDomainStmt) MarshalJSON() ([]byte, error) {
	type CreateDomainStmtMarshalAlias CreateDomainStmt
	return json.Marshal(map[string]interface{}{
		"CREATEDOMAINSTMT": (*CreateDomainStmtMarshalAlias)(&node),
	})
}

func (node *CreateDomainStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

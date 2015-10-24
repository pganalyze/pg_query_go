// Auto-generated - DO NOT EDIT

package pg_query

type CreateDomainStmt struct {
	Domainname  []Node         `json:"domainname"`  /* qualified name (list of Value strings) */
	TypeName    *TypeName      `json:"typeName"`    /* the base type */
	CollClause  *CollateClause `json:"collClause"`  /* untransformed COLLATE spec, if any */
	Constraints []Node         `json:"constraints"` /* constraints (list of Constraint nodes) */
}

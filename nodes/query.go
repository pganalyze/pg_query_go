// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Query struct {
	CommandType CmdType `json:"commandType"` /* select|insert|update|delete|utility */

	QuerySource QuerySource `json:"querySource"` /* where did I come from? */

	QueryId uint32 `json:"queryId"` /* query identifier (can be set by plugins) */

	CanSetTag bool `json:"canSetTag"` /* do I set the command result tag? */

	UtilityStmt Node `json:"utilityStmt"` /* non-null if this is DECLARE CURSOR or a
	 * non-optimizable statement */

	ResultRelation int `json:"resultRelation"` /* rtable index of target relation for
	 * INSERT/UPDATE/DELETE; 0 for SELECT */

	HasAggs         bool `json:"hasAggs"`         /* has aggregates in tlist or havingQual */
	HasWindowFuncs  bool `json:"hasWindowFuncs"`  /* has window functions in tlist */
	HasSubLinks     bool `json:"hasSubLinks"`     /* has subquery SubLink */
	HasDistinctOn   bool `json:"hasDistinctOn"`   /* distinctClause is from DISTINCT ON */
	HasRecursive    bool `json:"hasRecursive"`    /* WITH RECURSIVE was specified */
	HasModifyingCte bool `json:"hasModifyingCTE"` /* has INSERT/UPDATE/DELETE in WITH */
	HasForUpdate    bool `json:"hasForUpdate"`    /* FOR [KEY] UPDATE/SHARE was specified */

	CteList []Node `json:"cteList"` /* WITH list (of CommonTableExpr's) */

	Rtable   []Node    `json:"rtable"`   /* list of range table entries */
	Jointree *FromExpr `json:"jointree"` /* table join tree (FROM and WHERE clauses) */

	TargetList []Node `json:"targetList"` /* target list (of TargetEntry) */

	WithCheckOptions []Node `json:"withCheckOptions"` /* a list of WithCheckOption's */

	ReturningList []Node `json:"returningList"` /* return-values list (of TargetEntry) */

	GroupClause []Node `json:"groupClause"` /* a list of SortGroupClause's */

	HavingQual Node `json:"havingQual"` /* qualifications applied to groups */

	WindowClause []Node `json:"windowClause"` /* a list of WindowClause's */

	DistinctClause []Node `json:"distinctClause"` /* a list of SortGroupClause's */

	SortClause []Node `json:"sortClause"` /* a list of SortGroupClause's */

	LimitOffset Node `json:"limitOffset"` /* # of result tuples to skip (int8 expr) */
	LimitCount  Node `json:"limitCount"`  /* # of result tuples to return (int8 expr) */

	RowMarks []Node `json:"rowMarks"` /* a list of RowMarkClause's */

	SetOperations Node `json:"setOperations"` /* set-operation tree if this is top level of
	 * a UNION/INTERSECT/EXCEPT query */

	ConstraintDeps []Node `json:"constraintDeps"` /* a list of pg_constraint OIDs that the query
	 * depends on to be semantically valid */
}

func (node Query) MarshalJSON() ([]byte, error) {
	type QueryMarshalAlias Query
	return json.Marshal(map[string]interface{}{
		"QUERY": (*QueryMarshalAlias)(&node),
	})
}

func (node *Query) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

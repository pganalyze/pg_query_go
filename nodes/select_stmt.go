package pg_query

import "encoding/json"

// Make first letter of all fields uppercase
// char* => *string (so we can NULL)
// enum => const

type SelectStmt struct {
  /*
   * These fields are used only in "leaf" SelectStmts.
   */
  DistinctClause []Node `json:"distinctClause"` /* NULL, list of DISTINCT ON exprs, or
                 * lcons(NIL,NIL) for all (SELECT DISTINCT) */
  IntoClause *IntoClause `json:"intoClause"`		/* target for SELECT INTO */
  TargetList []Node	`json:"targetList"`	/* the target list (of ResTarget) */
  FromClause []Node	`json:"fromClause"`	/* the FROM clause */
  WhereClause Node `json:"whereClause"`	/* WHERE qualification */
  GroupClause []Node `json:"groupClause"`	/* GROUP BY clauses */
  HavingClause Node `json:"havingClause"`	/* HAVING conditional-expression */
  WindowClause []Node `json:"windowClause"`	/* WINDOW window_name AS (...), ... */

  /*
   * In a "leaf" node representing a VALUES list, the above fields are all
   * null, and instead this field is set.  Note that the elements of the
   * sublists are just expressions, without ResTarget decoration. Also note
   * that a list element can be DEFAULT (represented as a SetToDefault
   * node), regardless of the context of the VALUES list. It's up to parse
   * analysis to reject that where not valid.
   */
  ValuesLists []Node `json:"valuesLists"`	/* untransformed list of expression lists */

  /*
   * These fields are used in both "leaf" SelectStmts and upper-level
   * SelectStmts.
   */
  SortClause []Node `json:"sortClause"`		/* sort clause (a list of SortBy's) */
  LimitOffset Node `json:"limitOffset"`	/* # of result tuples to skip */
  LimitCount Node	`json:"limitCount"`	/* # of result tuples to return */
  LockingClause []Node `json:"lockingClause"`	/* FOR UPDATE (list of LockingClause's) */
  WithClause *WithClause	`json:"withClause"`	/* WITH clause */

  /*
   * These fields are used only in upper-level SelectStmts.
   */
  Op SetOperation	`json:"op"`		/* type of set op */
  All bool `json:"all"`			/* ALL specified? */
  Larg *SelectStmt `json:"larg"`	/* left child */
  Rarg *SelectStmt `json:"rarg"`	/* right child */
  /* Eventually add fields for CORRESPONDING spec here */
}

func (selectStmt SelectStmt) MarshalJSON() ([]byte, error) {
  type SelectStmtMarshalAlias SelectStmt
  return json.Marshal(map[string]interface{}{
    "SELECT": (*SelectStmtMarshalAlias)(&selectStmt),
  });
}

func (selectStmt *SelectStmt) UnmarshalJSON(input []byte) error {
  return UnmarshalNodeFieldJSON(input, selectStmt);
}

func (selectStmt SelectStmt) Deparse() string {
  panic("Not Implemented")
}

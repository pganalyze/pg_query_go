// Auto-generated - DO NOT EDIT

package pg_query

type RangeTblEntry struct {
	Rtekind RTEKind `json:"rtekind"` /* see above */

	/*
	 * XXX the fields applicable to only some rte kinds should be merged into
	 * a union.  I didn't do this yet because the diffs would impact a lot of
	 * code that is being actively worked on.  FIXME someday.
	 */

	/*
	 * Fields valid for a plain relation RTE (else zero):
	 */
	Relid   Oid  `json:"relid"`   /* OID of the relation */
	Relkind byte `json:"relkind"` /* relation kind (see pg_class.relkind) */

	/*
	 * Fields valid for a subquery RTE (else NULL):
	 */
	Subquery        *Query `json:"subquery"`         /* the sub-query */
	SecurityBarrier bool   `json:"security_barrier"` /* is from security_barrier view? */

	/*
	 * Fields valid for a join RTE (else NULL/zero):
	 *
	 * joinaliasvars is a list of (usually) Vars corresponding to the columns
	 * of the join result.  An alias Var referencing column K of the join
	 * result can be replaced by the K'th element of joinaliasvars --- but to
	 * simplify the task of reverse-listing aliases correctly, we do not do
	 * that until planning time.  In detail: an element of joinaliasvars can
	 * be a Var of one of the join's input relations, or such a Var with an
	 * implicit coercion to the join's output column type, or a COALESCE
	 * expression containing the two input column Vars (possibly coerced).
	 * Within a Query loaded from a stored rule, it is also possible for
	 * joinaliasvars items to be null pointers, which are placeholders for
	 * (necessarily unreferenced) columns dropped since the rule was made.
	 * Also, once planning begins, joinaliasvars items can be almost anything,
	 * as a result of subquery-flattening substitutions.
	 */
	Jointype      JoinType `json:"jointype"`      /* type of join */
	Joinaliasvars []Node   `json:"joinaliasvars"` /* list of alias-var expansions */

	/*
	 * Fields valid for a function RTE (else NIL/zero):
	 *
	 * When funcordinality is true, the eref->colnames list includes an alias
	 * for the ordinality column.  The ordinality column is otherwise
	 * implicit, and must be accounted for "by hand" in places such as
	 * expandRTE().
	 */
	Functions      []Node `json:"functions"`      /* list of RangeTblFunction nodes */
	Funcordinality bool   `json:"funcordinality"` /* is this called WITH ORDINALITY? */

	/*
	 * Fields valid for a values RTE (else NIL):
	 */
	ValuesLists      []Node `json:"values_lists"`      /* list of expression lists */
	ValuesCollations []Node `json:"values_collations"` /* OID list of column collation OIDs */

	/*
	 * Fields valid for a CTE RTE (else NULL/zero):
	 */
	Ctename          *string `json:"ctename"`          /* name of the WITH list item */
	Ctelevelsup      Index   `json:"ctelevelsup"`      /* number of query levels up */
	SelfReference    bool    `json:"self_reference"`   /* is this a recursive self-reference? */
	Ctecoltypes      []Node  `json:"ctecoltypes"`      /* OID list of column type OIDs */
	Ctecoltypmods    []Node  `json:"ctecoltypmods"`    /* integer list of column typmods */
	Ctecolcollations []Node  `json:"ctecolcollations"` /* OID list of column collation OIDs */

	/*
	 * Fields valid in all RTEs:
	 */
	Alias         *Alias   `json:"alias"`         /* user-written alias clause, if any */
	Eref          *Alias   `json:"eref"`          /* expanded reference names */
	Lateral       bool     `json:"lateral"`       /* subquery, function, or values is LATERAL? */
	Inh           bool     `json:"inh"`           /* inheritance requested? */
	InFromCl      bool     `json:"inFromCl"`      /* present in FROM clause? */
	RequiredPerms AclMode  `json:"requiredPerms"` /* bitmask of required access permissions */
	CheckAsUser   Oid      `json:"checkAsUser"`   /* if valid, check access as this role */
	SelectedCols  []uint32 `json:"selectedCols"`  /* columns needing SELECT permission */
	ModifiedCols  []uint32 `json:"modifiedCols"`  /* columns needing INSERT/UPDATE permission */
	SecurityQuals []Node   `json:"securityQuals"` /* any security barrier quals to apply */
}

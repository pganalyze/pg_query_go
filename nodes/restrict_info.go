// Auto-generated - DO NOT EDIT

package pg_query

type RestrictInfo struct {
	Clause *Expr `json:"clause"` /* the represented clause of WHERE or JOIN */

	IsPushedDown bool `json:"is_pushed_down"` /* TRUE if clause was pushed down in level */

	OuterjoinDelayed bool `json:"outerjoin_delayed"` /* TRUE if delayed by lower outer join */

	CanJoin bool `json:"can_join"` /* see comment above */

	Pseudoconstant bool `json:"pseudoconstant"` /* see comment above */

	/* The set of relids (varnos) actually referenced in the clause: */
	ClauseRelids []uint32 `json:"clause_relids"`

	/* The set of relids required to evaluate the clause: */
	RequiredRelids []uint32 `json:"required_relids"`

	/* If an outer-join clause, the outer-side relations, else NULL: */
	OuterRelids []uint32 `json:"outer_relids"`

	/* The relids used in the clause that are nullable by lower outer joins: */
	NullableRelids []uint32 `json:"nullable_relids"`

	/* These fields are set for any binary opclause: */
	LeftRelids  []uint32 `json:"left_relids"`  /* relids in left side of clause */
	RightRelids []uint32 `json:"right_relids"` /* relids in right side of clause */

	/* This field is NULL unless clause is an OR clause: */
	Orclause *Expr `json:"orclause"` /* modified clause with RestrictInfos */

	/* This field is NULL unless clause is potentially redundant: */
	ParentEc *EquivalenceClass `json:"parent_ec"` /* generating EquivalenceClass */

	/* cache space for cost and selectivity */
	EvalCost  QualCost    `json:"eval_cost"`  /* eval cost of clause; -1 if not yet set */
	NormSelec Selectivity `json:"norm_selec"` /* selectivity for "normal" (JOIN_INNER)
	 * semantics; -1 if not yet set; >1 means a
	 * redundant clause */
	OuterSelec Selectivity `json:"outer_selec"` /* selectivity for outer join semantics; -1 if
	 * not yet set */

	/* valid if clause is mergejoinable, else NIL */
	Mergeopfamilies []Node `json:"mergeopfamilies"` /* opfamilies containing clause operator */

	/* cache space for mergeclause processing; NULL if not yet set */
	LeftEc       *EquivalenceClass  `json:"left_ec"`       /* EquivalenceClass containing lefthand */
	RightEc      *EquivalenceClass  `json:"right_ec"`      /* EquivalenceClass containing righthand */
	LeftEm       *EquivalenceMember `json:"left_em"`       /* EquivalenceMember for lefthand */
	RightEm      *EquivalenceMember `json:"right_em"`      /* EquivalenceMember for righthand */
	ScanselCache []Node             `json:"scansel_cache"` /* list of MergeScanSelCache structs */

	/* transient workspace for use while considering a specific join path */
	OuterIsLeft bool `json:"outer_is_left"` /* T = outer var on left, F = on right */

	/* valid if clause is hashjoinable, else InvalidOid: */
	Hashjoinoperator Oid `json:"hashjoinoperator"` /* copy of clause operator */

	/* cache space for hashclause processing; -1 if not yet set */
	LeftBucketsize  Selectivity `json:"left_bucketsize"`  /* avg bucketsize of left side */
	RightBucketsize Selectivity `json:"right_bucketsize"` /* avg bucketsize of right side */
}

// Auto-generated - DO NOT EDIT

package pg_query

type MergeScanSelCache struct {
	/* Ordering details (cache lookup key) */
	Opfamily   Oid  `json:"opfamily"`    /* btree opfamily defining the ordering */
	Collation  Oid  `json:"collation"`   /* collation for the ordering */
	Strategy   int  `json:"strategy"`    /* sort direction (ASC or DESC) */
	NullsFirst bool `json:"nulls_first"` /* do NULLs come before normal values? */
	/* Results */
	Leftstartsel  Selectivity `json:"leftstartsel"`  /* first-join fraction for clause left side */
	Leftendsel    Selectivity `json:"leftendsel"`    /* last-join fraction for clause left side */
	Rightstartsel Selectivity `json:"rightstartsel"` /* first-join fraction for clause right side */
	Rightendsel   Selectivity `json:"rightendsel"`   /* last-join fraction for clause right side */
}

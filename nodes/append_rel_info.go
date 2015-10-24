// Auto-generated - DO NOT EDIT

package pg_query

type AppendRelInfo struct {
	/*
	 * These fields uniquely identify this append relationship.  There can be
	 * (in fact, always should be) multiple AppendRelInfos for the same
	 * parent_relid, but never more than one per child_relid, since a given
	 * RTE cannot be a child of more than one append parent.
	 */
	ParentRelid Index `json:"parent_relid"` /* RT index of append parent rel */
	ChildRelid  Index `json:"child_relid"`  /* RT index of append child rel */

	/*
	 * For an inheritance appendrel, the parent and child are both regular
	 * relations, and we store their rowtype OIDs here for use in translating
	 * whole-row Vars.  For a UNION-ALL appendrel, the parent and child are
	 * both subqueries with no named rowtype, and we store InvalidOid here.
	 */
	ParentReltype Oid `json:"parent_reltype"` /* OID of parent's composite type */
	ChildReltype  Oid `json:"child_reltype"`  /* OID of child's composite type */

	/*
	 * The N'th element of this list is a Var or expression representing the
	 * child column corresponding to the N'th column of the parent. This is
	 * used to translate Vars referencing the parent rel into references to
	 * the child.  A list element is NULL if it corresponds to a dropped
	 * column of the parent (this is only possible for inheritance cases, not
	 * UNION ALL).  The list elements are always simple Vars for inheritance
	 * cases, but can be arbitrary expressions in UNION ALL cases.
	 *
	 * Notice we only store entries for user columns (attno > 0).  Whole-row
	 * Vars are special-cased, and system columns (attno < 0) need no special
	 * translation since their attnos are the same for all tables.
	 *
	 * Caution: the Vars have varlevelsup = 0.  Be careful to adjust as needed
	 * when copying into a subquery.
	 */
	TranslatedVars []Node `json:"translated_vars"` /* Expressions in the child's Vars */

	/*
	 * We store the parent table's OID here for inheritance, or InvalidOid for
	 * UNION ALL.  This is only needed to help in generating error messages if
	 * an attempt is made to reference a dropped parent column.
	 */
	ParentReloid Oid `json:"parent_reloid"` /* OID of parent relation */
}

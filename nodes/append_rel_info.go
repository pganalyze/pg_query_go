// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * Append-relation info.
 *
 * When we expand an inheritable table or a UNION-ALL subselect into an
 * "append relation" (essentially, a list of child RTEs), we build an
 * AppendRelInfo for each child RTE.  The list of AppendRelInfos indicates
 * which child RTEs must be included when expanding the parent, and each
 * node carries information needed to translate Vars referencing the parent
 * into Vars referencing that child.
 *
 * These structs are kept in the PlannerInfo node's append_rel_list.
 * Note that we just throw all the structs into one list, and scan the
 * whole list when desiring to expand any one parent.  We could have used
 * a more complex data structure (eg, one list per parent), but this would
 * be harder to update during operations such as pulling up subqueries,
 * and not really any easier to scan.  Considering that typical queries
 * will not have many different append parents, it doesn't seem worthwhile
 * to complicate things.
 *
 * Note: after completion of the planner prep phase, any given RTE is an
 * append parent having entries in append_rel_list if and only if its
 * "inh" flag is set.  We clear "inh" for plain tables that turn out not
 * to have inheritance children, and (in an abuse of the original meaning
 * of the flag) we set "inh" for subquery RTEs that turn out to be
 * flattenable UNION ALL queries.  This lets us avoid useless searches
 * of append_rel_list.
 *
 * Note: the data structure assumes that append-rel members are single
 * baserels.  This is OK for inheritance, but it prevents us from pulling
 * up a UNION ALL member subquery if it contains a join.  While that could
 * be fixed with a more complex data structure, at present there's not much
 * point because no improvement in the plan could result.
 */
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

func (node AppendRelInfo) MarshalJSON() ([]byte, error) {
	type AppendRelInfoMarshalAlias AppendRelInfo
	return json.Marshal(map[string]interface{}{
		"APPENDRELINFO": (*AppendRelInfoMarshalAlias)(&node),
	})
}

func (node *AppendRelInfo) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["parent_relid"] != nil {
		err = json.Unmarshal(fields["parent_relid"], &node.ParentRelid)
		if err != nil {
			return
		}
	}

	if fields["child_relid"] != nil {
		err = json.Unmarshal(fields["child_relid"], &node.ChildRelid)
		if err != nil {
			return
		}
	}

	if fields["parent_reltype"] != nil {
		err = json.Unmarshal(fields["parent_reltype"], &node.ParentReltype)
		if err != nil {
			return
		}
	}

	if fields["child_reltype"] != nil {
		err = json.Unmarshal(fields["child_reltype"], &node.ChildReltype)
		if err != nil {
			return
		}
	}

	if fields["translated_vars"] != nil {
		node.TranslatedVars, err = UnmarshalNodeArrayJSON(fields["translated_vars"])
		if err != nil {
			return
		}
	}

	if fields["parent_reloid"] != nil {
		err = json.Unmarshal(fields["parent_reloid"], &node.ParentReloid)
		if err != nil {
			return
		}
	}

	return
}

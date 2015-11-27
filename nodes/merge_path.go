// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * A mergejoin path has these fields.
 *
 * Unlike other path types, a MergePath node doesn't represent just a single
 * run-time plan node: it can represent up to four.  Aside from the MergeJoin
 * node itself, there can be a Sort node for the outer input, a Sort node
 * for the inner input, and/or a Material node for the inner input.  We could
 * represent these nodes by separate path nodes, but considering how many
 * different merge paths are investigated during a complex join problem,
 * it seems better to avoid unnecessary palloc overhead.
 *
 * path_mergeclauses lists the clauses (in the form of RestrictInfos)
 * that will be used in the merge.
 *
 * Note that the mergeclauses are a subset of the parent relation's
 * restriction-clause list.  Any join clauses that are not mergejoinable
 * appear only in the parent's restrict list, and must be checked by a
 * qpqual at execution time.
 *
 * outersortkeys (resp. innersortkeys) is NIL if the outer path
 * (resp. inner path) is already ordered appropriately for the
 * mergejoin.  If it is not NIL then it is a PathKeys list describing
 * the ordering that must be created by an explicit Sort node.
 *
 * materialize_inner is TRUE if a Material node should be placed atop the
 * inner input.  This may appear with or without an inner Sort step.
 */
type MergePath struct {
	Jpath            JoinPath `json:"jpath"`
	PathMergeclauses []Node   `json:"path_mergeclauses"` /* join clauses to be used for merge */
	Outersortkeys    []Node   `json:"outersortkeys"`     /* keys for explicit sort, if any */
	Innersortkeys    []Node   `json:"innersortkeys"`     /* keys for explicit sort, if any */
	MaterializeInner bool     `json:"materialize_inner"` /* add Materialize to inner? */
}

func (node MergePath) MarshalJSON() ([]byte, error) {
	type MergePathMarshalAlias MergePath
	return json.Marshal(map[string]interface{}{
		"MERGEPATH": (*MergePathMarshalAlias)(&node),
	})
}

func (node *MergePath) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["jpath"] != nil {
		err = json.Unmarshal(fields["jpath"], &node.Jpath)
		if err != nil {
			return
		}
	}

	if fields["path_mergeclauses"] != nil {
		node.PathMergeclauses, err = UnmarshalNodeArrayJSON(fields["path_mergeclauses"])
		if err != nil {
			return
		}
	}

	if fields["outersortkeys"] != nil {
		node.Outersortkeys, err = UnmarshalNodeArrayJSON(fields["outersortkeys"])
		if err != nil {
			return
		}
	}

	if fields["innersortkeys"] != nil {
		node.Innersortkeys, err = UnmarshalNodeArrayJSON(fields["innersortkeys"])
		if err != nil {
			return
		}
	}

	if fields["materialize_inner"] != nil {
		err = json.Unmarshal(fields["materialize_inner"], &node.MaterializeInner)
		if err != nil {
			return
		}
	}

	return
}

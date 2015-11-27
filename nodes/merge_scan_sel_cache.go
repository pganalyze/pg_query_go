// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node MergeScanSelCache) MarshalJSON() ([]byte, error) {
	type MergeScanSelCacheMarshalAlias MergeScanSelCache
	return json.Marshal(map[string]interface{}{
		"MERGESCANSELCACHE": (*MergeScanSelCacheMarshalAlias)(&node),
	})
}

func (node *MergeScanSelCache) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["opfamily"] != nil {
		err = json.Unmarshal(fields["opfamily"], &node.Opfamily)
		if err != nil {
			return
		}
	}

	if fields["collation"] != nil {
		err = json.Unmarshal(fields["collation"], &node.Collation)
		if err != nil {
			return
		}
	}

	if fields["strategy"] != nil {
		err = json.Unmarshal(fields["strategy"], &node.Strategy)
		if err != nil {
			return
		}
	}

	if fields["nulls_first"] != nil {
		err = json.Unmarshal(fields["nulls_first"], &node.NullsFirst)
		if err != nil {
			return
		}
	}

	if fields["leftstartsel"] != nil {
		err = json.Unmarshal(fields["leftstartsel"], &node.Leftstartsel)
		if err != nil {
			return
		}
	}

	if fields["leftendsel"] != nil {
		err = json.Unmarshal(fields["leftendsel"], &node.Leftendsel)
		if err != nil {
			return
		}
	}

	if fields["rightstartsel"] != nil {
		err = json.Unmarshal(fields["rightstartsel"], &node.Rightstartsel)
		if err != nil {
			return
		}
	}

	if fields["rightendsel"] != nil {
		err = json.Unmarshal(fields["rightendsel"], &node.Rightendsel)
		if err != nil {
			return
		}
	}

	return
}

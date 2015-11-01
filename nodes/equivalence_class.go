// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type EquivalenceClass struct {
	EcOpfamilies []Node   `json:"ec_opfamilies"` /* btree operator family OIDs */
	EcCollation  Oid      `json:"ec_collation"`  /* collation, if datatypes are collatable */
	EcMembers    []Node   `json:"ec_members"`    /* list of EquivalenceMembers */
	EcSources    []Node   `json:"ec_sources"`    /* list of generating RestrictInfos */
	EcDerives    []Node   `json:"ec_derives"`    /* list of derived RestrictInfos */
	EcRelids     []uint32 `json:"ec_relids"`     /* all relids appearing in ec_members, except
	 * for child members (see below) */
	EcHasConst       bool              `json:"ec_has_const"`        /* any pseudoconstants in ec_members? */
	EcHasVolatile    bool              `json:"ec_has_volatile"`     /* the (sole) member is a volatile expr */
	EcBelowOuterJoin bool              `json:"ec_below_outer_join"` /* equivalence applies below an OJ */
	EcBroken         bool              `json:"ec_broken"`           /* failed to generate needed clauses? */
	EcSortref        Index             `json:"ec_sortref"`          /* originating sortclause label, or 0 */
	EcMerged         *EquivalenceClass `json:"ec_merged"`           /* set if merged into another EC */
}

func (node EquivalenceClass) MarshalJSON() ([]byte, error) {
	type EquivalenceClassMarshalAlias EquivalenceClass
	return json.Marshal(map[string]interface{}{
		"EQUIVALENCECLASS": (*EquivalenceClassMarshalAlias)(&node),
	})
}

func (node *EquivalenceClass) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

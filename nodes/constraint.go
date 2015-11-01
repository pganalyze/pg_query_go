// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Constraint struct {
	Contype ConstrType `json:"contype"` /* see above */

	/* Fields used for most/all constraint types: */
	Conname      *string `json:"conname"`      /* Constraint name, or NULL if unnamed */
	Deferrable   bool    `json:"deferrable"`   /* DEFERRABLE? */
	Initdeferred bool    `json:"initdeferred"` /* INITIALLY DEFERRED? */
	Location     int     `json:"location"`     /* token location, or -1 if unknown */

	/* Fields used for constraints with expressions (CHECK and DEFAULT): */
	IsNoInherit bool    `json:"is_no_inherit"` /* is constraint non-inheritable? */
	RawExpr     Node    `json:"raw_expr"`      /* expr, as untransformed parse tree */
	CookedExpr  *string `json:"cooked_expr"`   /* expr, as nodeToString representation */

	/* Fields used for unique constraints (UNIQUE and PRIMARY KEY): */
	Keys []Node `json:"keys"` /* String nodes naming referenced column(s) */

	/* Fields used for EXCLUSION constraints: */
	Exclusions []Node `json:"exclusions"` /* list of (IndexElem, operator name) pairs */

	/* Fields used for index constraints (UNIQUE, PRIMARY KEY, EXCLUSION): */
	Options    []Node  `json:"options"`    /* options from WITH clause */
	Indexname  *string `json:"indexname"`  /* existing index to use; otherwise NULL */
	Indexspace *string `json:"indexspace"` /* index tablespace; NULL for default */
	/* These could be, but currently are not, used for UNIQUE/PKEY: */
	AccessMethod *string `json:"access_method"` /* index access method; NULL for default */
	WhereClause  Node    `json:"where_clause"`  /* partial index predicate */

	/* Fields used for FOREIGN KEY constraints: */
	Pktable       *RangeVar `json:"pktable"`         /* Primary key table */
	FkAttrs       []Node    `json:"fk_attrs"`        /* Attributes of foreign key */
	PkAttrs       []Node    `json:"pk_attrs"`        /* Corresponding attrs in PK table */
	FkMatchtype   byte      `json:"fk_matchtype"`    /* FULL, PARTIAL, SIMPLE */
	FkUpdAction   byte      `json:"fk_upd_action"`   /* ON UPDATE action */
	FkDelAction   byte      `json:"fk_del_action"`   /* ON DELETE action */
	OldConpfeqop  []Node    `json:"old_conpfeqop"`   /* pg_constraint.conpfeqop of my former self */
	OldPktableOid Oid       `json:"old_pktable_oid"` /* pg_constraint.confrelid of my former self */

	/* Fields used for constraints that allow a NOT VALID specification */
	SkipValidation bool `json:"skip_validation"` /* skip validation of existing rows? */
	InitiallyValid bool `json:"initially_valid"` /* mark the new constraint as valid? */
}

func (node Constraint) MarshalJSON() ([]byte, error) {
	type ConstraintMarshalAlias Constraint
	return json.Marshal(map[string]interface{}{
		"CONSTRAINT": (*ConstraintMarshalAlias)(&node),
	})
}

func (node *Constraint) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

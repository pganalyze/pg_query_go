// Auto-generated - DO NOT EDIT

package pg_query

type CreateTrigStmt struct {
	Trigname *string   `json:"trigname"` /* TRIGGER's name */
	Relation *RangeVar `json:"relation"` /* relation trigger is on */
	Funcname []Node    `json:"funcname"` /* qual. name of function to call */
	Args     []Node    `json:"args"`     /* list of (T_String) Values or NIL */
	Row      bool      `json:"row"`      /* ROW/STATEMENT */
	/* timing uses the TRIGGER_TYPE bits defined in catalog/pg_trigger.h */
	Timing int16 `json:"timing"` /* BEFORE, AFTER, or INSTEAD */
	/* events uses the TRIGGER_TYPE bits defined in catalog/pg_trigger.h */
	Events       int16  `json:"events"`       /* "OR" of INSERT/UPDATE/DELETE/TRUNCATE */
	Columns      []Node `json:"columns"`      /* column names, or NIL for all columns */
	WhenClause   Node   `json:"whenClause"`   /* qual expression, or NULL if none */
	Isconstraint bool   `json:"isconstraint"` /* This is a constraint trigger */
	/* The remaining fields are only used for constraint triggers */
	Deferrable   bool      `json:"deferrable"`   /* [NOT] DEFERRABLE */
	Initdeferred bool      `json:"initdeferred"` /* INITIALLY {DEFERRED|IMMEDIATE} */
	Constrrel    *RangeVar `json:"constrrel"`    /* opposite relation, if RI trigger */
}

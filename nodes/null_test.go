// Auto-generated - DO NOT EDIT

package pg_query

type NullTest struct {
	Xpr          Expr         `json:"xpr"`
	Arg          *Expr        `json:"arg"`          /* input expression */
	Nulltesttype NullTestType `json:"nulltesttype"` /* IS NULL, IS NOT NULL */
	Argisrow     bool         `json:"argisrow"`     /* T if input is of a composite type */
}

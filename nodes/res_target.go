package pg_query

type ResTarget struct {
	Name string	`json:"name"`		/* column name or NULL */
	Indirection []Node `json:"indirection"` /* subscripts, field names, and '*', or NIL */
	Val Node `json:"val"`			/* the value expression to compute or assign */
	Location int `json:"location"`		/* token location, or -1 if unknown */
}

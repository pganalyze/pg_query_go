package pg_query

type A_Const struct {
	Val Value	`json:"val"`		/* value (includes type info, see value.h) */
	Location int `json:"location"`		/* token location, or -1 if unknown */
}

package pg_query

// TODO(LukasFittl): This might make more sense as separate types per embedded type, and Value as an interface

type Value struct {
	//ValueType NodeTag			/* tag appropriately (eg. T_String) */
	Ival int `json:"ival"`		/* machine integer */
	Str string `json:"str"`		/* string */
}

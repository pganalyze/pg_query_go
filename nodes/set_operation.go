package pg_query

type SetOperation uint

const (
	SETOP_NONE = iota
	SETOP_UNION
	SETOP_INTERSECT
	SETOP_EXCEPT
)

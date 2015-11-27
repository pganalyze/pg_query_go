package pg_query

type Node interface {
	Deparse() string
}

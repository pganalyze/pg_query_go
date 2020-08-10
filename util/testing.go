package util

import (
	nodes "github.com/lfittl/pg_query_go/nodes"
)

func MakeStrPtr(str string) *string {
	return &str
}

func MakeStrNode(str string) nodes.String {
	return nodes.String{Str: str}
}

func MakeIntNode(ival int64) *nodes.Node {
	return &nodes.Node{Node: &nodes.Node_Integer{Integer: &nodes.Integer{Ival: int32(ival)}}} // FIXME
}

func MakeListNode(items []*nodes.Node) nodes.List {
	return nodes.List{Items: items}
}

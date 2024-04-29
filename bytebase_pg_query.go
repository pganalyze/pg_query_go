package pg_query

import (
	"fmt"

	"github.com/pganalyze/pg_query_go/v5/parser"
	"google.golang.org/protobuf/proto"
)

type DeparseType int

const (
	DeparseTypeExpr DeparseType = iota
	DeparseTypeExclusion
	DeparseTypeDataType
)

func DeparseNode(tp DeparseType, node *Node) (output string, err error) {
	switch tp {
	case DeparseTypeExpr:
	case DeparseTypeDataType:
	default:
		return "", fmt.Errorf("deparse node failed: unsupported deparse type %d", tp)
	}

	tree := &ParseResult{
		Stmts: []*RawStmt{
			{
				Stmt: node,
			},
		},
	}
	protobufTree, err := proto.Marshal(tree)
	if err != nil {
		return "", err
	}
	return parser.DeparseNodeFromProtobuf(int(tp), protobufTree)
}

func DeparseNodes(tp DeparseType, nodes []*Node) (output string, err error) {
	switch tp {
	case DeparseTypeExclusion:
		tree := &ParseResult{}
		for _, node := range nodes {
			tree.Stmts = append(tree.Stmts, &RawStmt{
				Stmt: node,
			})
		}
		protobufTree, err := proto.Marshal(tree)
		if err != nil {
			return "", err
		}
		return parser.DeparseNodeFromProtobuf(int(tp), protobufTree)
	default:
		return "", fmt.Errorf("deparse node failed: unsupported deparse type %d", tp)
	}
}

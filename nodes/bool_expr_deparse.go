// Auto-generated - DO NOT EDIT

package pg_query

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func (node BoolExpr) Deparse(ctx DeparseContext) (string, error) {
	// There is no BOOL_EXPR_NOT in go for some reason?
	switch node.Boolop {
	case AND_EXPR:
		return node.deparseBoolExprAnd()
	case OR_EXPR:
		return node.deparseBoolExprOr()
	default:
		return "", errors.New(fmt.Sprintf("cannot handle bool expression type (%d)", node.Boolop))
	}
}

// I feel like these two functions could probably be merged.

func (node BoolExpr) deparseBoolExprAnd() (string, error) {
	if node.Args.Items == nil || len(node.Args.Items) == 0 {
		return "", errors.New("args cannot be empty for boolean expression")
	}
	args := make([]string, len(node.Args.Items))
	for i, arg := range node.Args.Items {
		if str, err := arg.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			t := reflect.TypeOf(arg)
			if t == reflect.TypeOf(BoolExpr{}) && arg.(BoolExpr).Boolop == OR_EXPR {
				args[i] = fmt.Sprintf("(%s)", str)
			} else {
				args[i] = str
			}
		}
	}
	return strings.Join(args, " AND "), nil
}

func (node BoolExpr) deparseBoolExprOr() (string, error) {
	if node.Args.Items == nil || len(node.Args.Items) == 0 {
		return "", errors.New("args cannot be empty for boolean expression")
	}
	args := make([]string, len(node.Args.Items))
	for i, arg := range node.Args.Items {
		if str, err := arg.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			t := reflect.TypeOf(arg)
			if t == reflect.TypeOf(BoolExpr{}) && (arg.(BoolExpr).Boolop == OR_EXPR || arg.(BoolExpr).Boolop == AND_EXPR) {
				args[i] = fmt.Sprintf("(%s)", str)
			} else {
				args[i] = str
			}
		}
	}
	return strings.Join(args, " OR "), nil
}

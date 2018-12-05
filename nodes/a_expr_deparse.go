// Auto-generated - DO NOT EDIT

package pg_query

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func (node A_Expr) Deparse(ctx DeparseContext) (string, error) {
	switch node.Kind {
	case AEXPR_OP:
		return node.deparseAExpr(ctx)
	case AEXPR_OP_ANY:
		return node.deparseAExprAny(ctx)
	case AEXPR_IN:
		return node.deparseAExprIn(ctx)
	case AEXPR_BETWEEN, AEXPR_NOT_BETWEEN, AEXPR_BETWEEN_SYM, AEXPR_NOT_BETWEEN_SYM:
		return node.deparseAExprBetween(ctx)
	case AEXPR_NULLIF:
		return node.deparseAExprNullIf(ctx)
	default:
		return "", errors.New(fmt.Sprintf("could not parse AExpr of kind: %d, not implemented", node.Kind))
	}
}

func (node A_Expr) deparseAExpr(ctx DeparseContext) (string, error) {
	out := make([]string, 0)

	if node.Lexpr == nil {
		return "", errors.New("lexpr of expression cannot be null")
	}

	if node.Lexpr == nil {
		return "", errors.New("rexpr of expression cannot be null")
	}

	if node.Name.Items == nil || len(node.Name.Items) == 0 {
		return "", errors.New("error, expression name cannot be null")
	}

	switch n := node.Lexpr.(type) {
	case List:
		if n.Items == nil || len(n.Items) == 0 {
			return "", errors.New("lexpr list cannot be empty")
		}
		if str, err := n.Items[0].Deparse(ctx); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	default:
		if str, err := n.Deparse(ctx); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	}

	if str, err := node.Rexpr.Deparse(ctx); err != nil {
		return "", err
	} else {
		out = append(out, str)
	}

	if name, err := node.Name.Items[0].Deparse(DeparseContextOperator); err != nil {
		return "", err
	} else {
		result := strings.Join(out, fmt.Sprintf(" %s ", name))
		if ctx != DeparseContextNone {
			result = fmt.Sprintf("(%s)", result)
		}
		return result, nil
	}
}

func (node A_Expr) deparseAExprIn(ctx DeparseContext) (string, error) {
	out := make([]string, 0)

	if node.Rexpr == nil {
		return "", errors.New("rexpr of IN expression cannot be null")
	}

	if strs, err := node.Rexpr.(List).DeparseList(DeparseContextNone); err != nil {
		return "", err
	} else {
		out = append(out, strs...)
	}

	if node.Name.Items == nil || len(node.Name.Items) == 0 {
		return "", errors.New("names of IN expression cannot be empty")
	}

	if strs, err := node.Name.DeparseList(DeparseContextOperator); err != nil {
		return "", err
	} else {
		operator := ""
		if reflect.DeepEqual(strs, []string{"="}) {
			operator = "IN"
		} else {
			operator = "NOT IN"
		}

		if node.Lexpr == nil {
			return "", errors.New("lexpr of IN expression cannot be null")
		}

		if str, err := node.Lexpr.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			return fmt.Sprintf("%s %s (%s)", str, operator, strings.Join(out, ", ")), nil
		}
	}
}

func (node A_Expr) deparseAExprAny(ctx DeparseContext) (string, error) {
	out := make([]string, 0)
	if str, err := node.Lexpr.Deparse(DeparseContextNone); err != nil {
		return "", err
	} else {
		out = append(out, str)
	}

	if str, err := node.Rexpr.Deparse(DeparseContextNone); err != nil {
		return "", err
	} else {
		out = append(out, fmt.Sprintf("ANY(%s)", str))
	}

	// TODO (elliotcourant) this seems a bit weird that we are just taking the first item for this. Revist this in the future?
	if str, err := node.Name.Items[0].Deparse(DeparseContextOperator); err != nil {
		return "", err
	} else {
		return strings.Join(out, str), nil
	}
}

func (node A_Expr) deparseAExprBetween(ctx DeparseContext) (string, error) {
	between := ""
	switch node.Kind {
	case AEXPR_BETWEEN:
		between = "BETWEEN"
	case AEXPR_NOT_BETWEEN:
		between = "NOT BETWEEN"
	case AEXPR_BETWEEN_SYM:
		between = "BETWEEN SYMMETRIC"
	case AEXPR_NOT_BETWEEN_SYM:
		between = "NOT BETWEEN SYMMETRIC"
	}

	name, err := node.Lexpr.Deparse(DeparseContextNone)
	if err != nil {
		return "", err
	}

	rightExpression := node.Rexpr.(List)
	out := make([]string, len(rightExpression.Items))
	for i, expr := range rightExpression.Items {
		if str, err := expr.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out[i] = str
		}
	}

	return fmt.Sprintf("%s %s %s", name, between, strings.Join(out, " AND ")), nil
}

func (node A_Expr) deparseAExprNullIf(ctx DeparseContext) (string, error) {
	leftString, err := node.Lexpr.Deparse(DeparseContextNone)
	if err != nil {
		return "", err
	}

	rightString, err := node.Rexpr.Deparse(DeparseContextNone)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("NULLIF(%s, %s)", leftString, rightString), nil
}

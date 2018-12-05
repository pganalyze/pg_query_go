// Auto-generated - DO NOT EDIT

package pg_query

import (
	"errors"
	"fmt"
	"strings"
)

func (node JoinExpr) Deparse(ctx DeparseContext) (string, error) {
	out := make([]string, 0)

	if node.Larg == nil {
		return "", errors.New("larg of join cannot be null")
	}

	if str, err := node.Larg.Deparse(DeparseContextNone); err != nil {
		return "", err
	} else {
		out = append(out, str)
	}

	switch node.Jointype {
	case JOIN_INNER:
		if node.IsNatural {
			out = append(out, "NATURAL")
		} else if node.Quals == nil && (node.UsingClause.Items == nil || len(node.UsingClause.Items) == 0) {
			out = append(out, "CROSS")
		}
	case JOIN_LEFT:
		out = append(out, "LEFT")
	case JOIN_FULL:
		out = append(out, "FULL")
	case JOIN_RIGHT:
		out = append(out, "RIGHT")
	default:
		return "", errors.New(fmt.Sprintf("cannot handle join type (%d)", node.Jointype))
	}
	out = append(out, "JOIN")

	if node.Rarg == nil {
		return "", errors.New("rarg of join cannot be null")
	}

	if str, err := node.Rarg.Deparse(DeparseContextNone); err != nil {
		return "", err
	} else {
		out = append(out, str)
	}

	if node.Quals != nil {
		out = append(out, "ON")
		if str, err := node.Quals.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	}

	if node.UsingClause.Items != nil && len(node.UsingClause.Items) > 0 {
		clauses := make([]string, len(node.UsingClause.Items))
		for i, field := range node.UsingClause.Items {
			if str, err := field.Deparse(DeparseContextSelect); err != nil {
				return "", err
			} else {
				clauses[i] = str
			}
		}
		out = append(out, fmt.Sprintf("USING (%s)", strings.Join(clauses, ", ")))
	}

	return strings.Join(out, " "), nil
}

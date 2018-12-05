// Auto-generated - DO NOT EDIT

package pg_query

import (
	"fmt"
	"strings"
)

func (node Constraint) Deparse(ctx DeparseContext) (string, error) {
	out := make([]string, 0)
	if node.Conname != nil {
		out = append(out, "CONSTRAINT")
		out = append(out, *node.Conname)
	}
	switch node.Contype {
	case CONSTR_NULL:
		out = append(out, "NULL")
	case CONSTR_NOTNULL:
		out = append(out, "NOT NULL")
	case CONSTR_DEFAULT:
		out = append(out, "DEFAULT")
	case CONSTR_CHECK:
		out = append(out, "CHECK")
	case CONSTR_PRIMARY:
		out = append(out, "PRIMARY KEY")
	case CONSTR_UNIQUE:
		out = append(out, "UNIQUE")
	case CONSTR_EXCLUSION:
		out = append(out, "EXCLUSION")
	case CONSTR_FOREIGN:
		out = append(out, "FOREIGN KEY")
	}

	if node.RawExpr != nil {
		if expr, err := node.RawExpr.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			if aexpr, ok := node.RawExpr.(A_Expr); ok && aexpr.Kind == AEXPR_OP {
				out = append(out, fmt.Sprintf("(%s)", expr))
			} else {
				out = append(out, expr)
			}
		}
	}

	if node.Keys.Items != nil && len(node.Keys.Items) > 0 {
		if list, err := node.Keys.DeparseList(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, fmt.Sprintf("(%s)", strings.Join(list, ", ")))
		}
	}

	if node.FkAttrs.Items != nil && len(node.FkAttrs.Items) > 0 {
		if list, err := node.FkAttrs.DeparseList(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, fmt.Sprintf("(%s)", strings.Join(list, ", ")))
		}
	}

	if node.Pktable != nil {
		if list, err := node.PkAttrs.DeparseList(DeparseContextNone); err != nil {
			return "", err
		} else {
			if pk, err := node.Pktable.Deparse(DeparseContextNone); err != nil {
				return "", err
			} else {
				out = append(out, fmt.Sprintf("REFERENCES %s (%s)", pk, strings.Join(list, ", ")))
			}
		}
	}

	if node.SkipValidation {
		out = append(out, "NOT VALID")
	}

	if node.Indexname != nil {
		out = append(out, fmt.Sprintf("USING INDEX %s", *node.Indexname))
	}
	return strings.Join(out, " "), nil
}
